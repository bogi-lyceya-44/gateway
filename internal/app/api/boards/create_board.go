package boards

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/api/mappers"
	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/boards"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Api) CreateBoard(
	ctx context.Context,
	request *desc.CreateBoardRequest,
) (*desc.CreateBoardResponse, error) {
	createdBoard, err := a.boardService.CreateBoard(
		ctx,
		mappers.MapCreateBoardPrototypeToDomain(request.GetBoardToCreate()),
	)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return nil, status.Error(st.Code(), st.Message())
		}

		return nil, status.Errorf(codes.Internal, "creating board: %v", err)
	}

	return &desc.CreateBoardResponse{Id: createdBoard}, nil
}
