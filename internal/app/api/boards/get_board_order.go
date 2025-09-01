package boards

import (
	"context"

	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/boards"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Api) GetBoardsOrder(
	ctx context.Context,
	request *desc.GetBoardsOrderRequest,
) (*desc.GetBoardsOrderResponse, error) {
	order, err := a.boardService.GetBoardOrder(
		ctx,
	)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return nil, status.Error(st.Code(), st.Message())
		}

		return nil, status.Errorf(codes.Internal, "getting board order: %v", err)
	}

	return &desc.GetBoardsOrderResponse{Order: order}, nil
}
