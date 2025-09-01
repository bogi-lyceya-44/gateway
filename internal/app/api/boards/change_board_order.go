package boards

import (
	"context"

	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/boards"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Api) ChangeBoardOrder(
	ctx context.Context,
	request *desc.ChangeBoardOrderRequest,
) (*desc.ChangeBoardOrderResponse, error) {
	err := a.boardService.ChangeBoardOrder(
		ctx,
		request.GetId(),
		request.GetOrderToChange().GetPlace(),
	)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return nil, status.Error(st.Code(), st.Message())
		}

		return nil, status.Errorf(codes.Internal, "changing order: %v", err)
	}

	return &desc.ChangeBoardOrderResponse{}, nil
}
