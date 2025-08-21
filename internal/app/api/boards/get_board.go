package boards

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/api/mappers"
	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/boards"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Api) GetBoard(
	ctx context.Context,
	request *desc.GetBoardRequest,
) (*desc.GetBoardResponse, error) {
	board, err := a.boardService.GetBoard(
		ctx,
		request.GetId(),
	)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return nil, status.Error(st.Code(), st.Message())
		}

		return nil, status.Errorf(codes.Internal, "getting board: %v", err)
	}

	mappedBoard := mappers.MapDomainBoardToProto(
		board,
	)

	return &desc.GetBoardResponse{Board: mappedBoard}, nil
}
