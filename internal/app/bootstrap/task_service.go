package bootstrap

import (
	"context"

	tasks_api "github.com/bogi-lyceya-44/gateway/internal/app/api/tasks"
	tasks_repo "github.com/bogi-lyceya-44/gateway/internal/app/repositories/tasks"
	tasks_service "github.com/bogi-lyceya-44/gateway/internal/app/services/tasks"
	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/tasks"
	desc_vendor "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/tasks"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitTaskService(
	ctx context.Context,
	clientConn grpc.ClientConnInterface,
	app *App,
) error {
	taskServiceClient := desc_vendor.NewTaskServiceClient(clientConn)
	taskRepository := tasks_repo.New(taskServiceClient)

	taskService := tasks_service.New(taskRepository)

	taskApi := tasks_api.New(taskService)

	desc.RegisterTaskServiceServer(app.grpcServer, taskApi)

	if err := desc.RegisterTaskServiceHandlerFromEndpoint(
		ctx,
		app.mux,
		app.grpcAddr,
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	); err != nil {
		return errors.Wrap(err, "registering task service gateway")
	}

	return nil
}
