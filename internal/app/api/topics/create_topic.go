package topics

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/api/mappers"
	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/topics"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Api) CreateTopic(
	ctx context.Context,
	request *desc.CreateTopicRequest,
) (*desc.CreateTopicResponse, error) {
	mappedTopic := mappers.MapCreateTopicPrototypeToDomain(
		request.GetTopicToCreate(),
	)

	createdTopic, err := a.topicService.CreateTopic(
		ctx,
		mappedTopic,
	)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return nil, status.Error(st.Code(), st.Message())
		}

		return nil, status.Errorf(codes.Internal, "creating topic: %v", err)
	}

	return &desc.CreateTopicResponse{Id: createdTopic}, nil
}
