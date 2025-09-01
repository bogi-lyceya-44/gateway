package boards

import (
	"context"

	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/boards"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Api) DeleteBoard(
	ctx context.Context,
	request *desc.DeleteBoardRequest,
) (*desc.DeleteBoardResponse, error) {
	err := a.boardService.DeleteBoard(
		ctx,
		request.GetId(),
	)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return nil, status.Error(st.Code(), st.Message())
		}

		return nil, status.Errorf(codes.Internal, "deleting board: %v", err)
	}

	return &desc.DeleteBoardResponse{}, nil
}
