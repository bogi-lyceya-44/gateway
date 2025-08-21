package boards

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/gateway/internal/app/api/mappers"
	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/boards"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Api) GetBoardContent(
	ctx context.Context,
	request *desc.GetBoardContentRequest,
) (*desc.GetBoardContentResponse, error) {
	content, err := a.boardService.GetBoardContent(
		ctx,
		request.GetId(),
	)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return nil, status.Error(st.Code(), st.Message())
		}

		return nil, status.Errorf(codes.Internal, "getting board content: %v", err)
	}

	mappedContent, err := utils.MapWithError(
		content,
		mappers.MapDomainTopicWithTaskToProto,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "mapping topics: %v", err)
	}

	return &desc.GetBoardContentResponse{
		Topics: mappedContent,
	}, nil
}
