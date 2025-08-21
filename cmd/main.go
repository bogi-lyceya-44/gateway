package main

import (
	"log"
	"net"

	"github.com/bogi-lyceya-44/gateway/config"
	"github.com/bogi-lyceya-44/gateway/internal/app/bootstrap"
	"github.com/pkg/errors"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(errors.Wrap(err, "init config"))
	}

	bootstrap.InitCloser()

	ctx, err := bootstrap.InitGlobalContext()
	if err != nil {
		log.Fatal(errors.Wrap(err, "init global context"))
	}

	clientConn, err := bootstrap.InitClientConnection(
		net.JoinHostPort(cfg.TaskTracker.Host, cfg.TaskTracker.Port),
	)
	if err != nil {
		log.Fatal(errors.Wrap(err, "init client connection"))
	}

	app := bootstrap.InitApp(cfg)

	if err := bootstrap.InitBoardService(
		ctx,
		clientConn,
		app,
	); err != nil {
		log.Fatal(errors.Wrap(err, "init board service"))
	}

	if err := bootstrap.InitTaskService(
		ctx,
		clientConn,
		app,
	); err != nil {
		log.Fatal(errors.Wrap(err, "init task service"))
	}

	if err := bootstrap.InitTopicService(
		ctx,
		clientConn,
		app,
	); err != nil {
		log.Fatal(errors.Wrap(err, "init topic service"))
	}

	if err = app.Run(ctx); err != nil {
		log.Fatal(errors.Wrap(err, "running app"))
	}
}
