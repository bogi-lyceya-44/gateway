package topics

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/api/mappers"
	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/topics"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Api) GetTopic(
	ctx context.Context,
	request *desc.GetTopicRequest,
) (*desc.GetTopicResponse, error) {
	topic, err := a.topicService.GetTopic(
		ctx,
		request.GetId(),
	)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return nil, status.Error(st.Code(), st.Message())
		}

		return nil, status.Errorf(codes.Internal, "deleting topic: %v", err)
	}

	mappedTopic := mappers.MapDomainTopicToProto(topic)

	return &desc.GetTopicResponse{Topic: mappedTopic}, nil
}
