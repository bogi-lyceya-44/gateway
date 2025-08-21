package bootstrap

import (
	"context"

	topics_api "github.com/bogi-lyceya-44/gateway/internal/app/api/topics"
	topics_repo "github.com/bogi-lyceya-44/gateway/internal/app/repositories/topics"
	topics_service "github.com/bogi-lyceya-44/gateway/internal/app/services/topics"
	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/topics"
	desc_vendor "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/topics"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitTopicService(
	ctx context.Context,
	clientConn grpc.ClientConnInterface,
	app *App,
) error {
	topicServiceClient := desc_vendor.NewTopicServiceClient(clientConn)
	topicRepository := topics_repo.New(topicServiceClient)

	topicService := topics_service.New(topicRepository)

	topicApi := topics_api.New(topicService)

	desc.RegisterTopicServiceServer(app.grpcServer, topicApi)

	if err := desc.RegisterTopicServiceHandlerFromEndpoint(
		ctx,
		app.mux,
		app.grpcAddr,
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	); err != nil {
		return errors.Wrap(err, "registering topic service gateway")
	}

	return nil
}
