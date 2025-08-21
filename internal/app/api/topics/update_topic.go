package topics

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/api/mappers"
	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/topics"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Api) UpdateTopic(
	ctx context.Context,
	request *desc.UpdateTopicRequest,
) (*desc.UpdateTopicResponse, error) {
	mappedTopic := mappers.MapUpdateTopicPrototypeToDomain(
		request.GetId(),
		request.GetTopicToUpdate(),
	)

	err := a.topicService.UpdateTopic(
		ctx,
		mappedTopic,
	)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return nil, status.Error(st.Code(), st.Message())
		}

		return nil, status.Errorf(codes.Internal, "updating topic: %v", err)
	}

	return &desc.UpdateTopicResponse{}, nil
}
