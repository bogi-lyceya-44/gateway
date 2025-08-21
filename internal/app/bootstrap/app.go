package bootstrap

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/bogi-lyceya-44/common/pkg/closer"
	"github.com/bogi-lyceya-44/gateway/config"
	"github.com/bogi-lyceya-44/gateway/docs"
	"github.com/flowchartsman/swaggerui"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const ShutdownTimeoutS = 5

type App struct {
	grpcServer *grpc.Server
	mux        *runtime.ServeMux

	grpcAddr    string
	gatewayAddr string
}

func InitApp(
	cfg *config.Config,
) *App {
	grpcServer := grpc.NewServer()
	mux := runtime.NewServeMux()

	grpcAddr := net.JoinHostPort(cfg.GRPC.Host, cfg.GRPC.Port)
	gatewayAddr := net.JoinHostPort(cfg.Gateway.Host, cfg.Gateway.Port)

	reflection.Register(grpcServer)

	return &App{
		grpcServer:  grpcServer,
		mux:         mux,
		grpcAddr:    grpcAddr,
		gatewayAddr: gatewayAddr,
	}
}

func (a *App) Run(ctx context.Context) error {
	eg, _ := errgroup.WithContext(ctx)

	lis, err := net.Listen(
		"tcp",
		a.grpcAddr,
	)
	if err != nil {
		return errors.Wrap(err, "listener create")
	}

	if err = a.mux.HandlePath(
		"GET",
		"/docs",
		func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
			http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
		},
	); err != nil {
		return errors.Wrap(err, "registering swagger json")
	}

	if err = a.mux.HandlePath(
		"GET",
		"/swagger/{path=**}",
		func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
			http.StripPrefix("/swagger", swaggerui.Handler(docs.Spec)).ServeHTTP(w, r)
		},
	); err != nil {
		return errors.Wrap(err, "registering swagger json")
	}

	httpServer := http.Server{
		Addr:    a.gatewayAddr,
		Handler: a.mux,
	}

	eg.Go(func() error {
		if err := httpServer.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			return errors.Wrap(err, "listening http")
		}

		return nil
	})

	eg.Go(func() error {
		if err := a.grpcServer.Serve(lis); err != nil &&
			!errors.Is(err, grpc.ErrServerStopped) {
			return errors.Wrap(err, "listening grpc")
		}

		return nil
	})

	if err := closer.AddCallback(
		CloserAppGroup,
		func() error {
			defer lis.Close()

			a.grpcServer.GracefulStop()

			shutdownCtx, cancel := context.WithTimeout(
				context.Background(),
				ShutdownTimeoutS*time.Second,
			)
			defer cancel()

			if err := httpServer.Shutdown(shutdownCtx); err != nil {
				return errors.Wrap(err, "shutting down http")
			}

			return nil
		},
	); err != nil {
		return errors.Wrap(err, "app callback")
	}

	log.Printf("http server listening on %s", a.gatewayAddr)
	log.Printf("grpc server listening on %s", a.grpcAddr)

	err = eg.Wait()
	if err != nil {
		return errors.Wrap(err, "starting app")
	}

	err = closer.Wait()
	if err != nil {
		return errors.Wrap(err, "executing callbacks")
	}

	return nil
}
