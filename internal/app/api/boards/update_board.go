package boards

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/api/mappers"
	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/boards"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Api) UpdateBoard(
	ctx context.Context,
	request *desc.UpdateBoardRequest,
) (*desc.UpdateBoardResponse, error) {
	err := a.boardService.UpdateBoard(
		ctx,
		mappers.MapUpdatedBoardPrototypeToDomain(
			request.GetId(),
			request.GetBoardToUpdate(),
		),
	)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return nil, status.Error(st.Code(), st.Message())
		}

		return nil, status.Errorf(codes.Internal, "updating board: %v", err)
	}

	return &desc.UpdateBoardResponse{}, nil
}
