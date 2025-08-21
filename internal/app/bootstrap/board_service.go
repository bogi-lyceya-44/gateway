package bootstrap

import (
	"context"

	boards_api "github.com/bogi-lyceya-44/gateway/internal/app/api/boards"
	boards_repo "github.com/bogi-lyceya-44/gateway/internal/app/repositories/boards"
	boards_service "github.com/bogi-lyceya-44/gateway/internal/app/services/boards"
	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/boards"
	desc_vendor "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/boards"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitBoardService(
	ctx context.Context,
	clientConn grpc.ClientConnInterface,
	app *App,
) error {
	boardServiceClient := desc_vendor.NewBoardServiceClient(clientConn)
	boardRepository := boards_repo.New(boardServiceClient)

	boardService := boards_service.New(boardRepository)

	boardApi := boards_api.New(boardService)

	desc.RegisterBoardServiceServer(app.grpcServer, boardApi)

	if err := desc.RegisterBoardServiceHandlerFromEndpoint(
		ctx,
		app.mux,
		app.grpcAddr,
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	); err != nil {
		return errors.Wrap(err, "registering board service gateway")
	}

	return nil
}
