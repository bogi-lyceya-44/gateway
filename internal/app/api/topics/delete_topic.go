package topics

import (
	"context"

	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/topics"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Api) DeleteTopic(
	ctx context.Context,
	request *desc.DeleteTopicRequest,
) (*desc.DeleteTopicResponse, error) {
	err := a.topicService.DeleteTopic(
		ctx,
		request.GetId(),
	)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return nil, status.Error(st.Code(), st.Message())
		}

		return nil, status.Errorf(codes.Internal, "deleting topic: %v", err)
	}

	return &desc.DeleteTopicResponse{}, nil
}
