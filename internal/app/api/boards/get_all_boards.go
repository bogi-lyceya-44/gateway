package boards

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/gateway/internal/app/api/mappers"
	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/boards"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Api) GetAllBoards(
	ctx context.Context,
	request *desc.GetAllBoardsRequest,
) (*desc.GetAllBoardsResponse, error) {
	allBoards, err := a.boardService.GetAllBoards(
		ctx,
	)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return nil, status.Error(st.Code(), st.Message())
		}

		return nil, status.Errorf(codes.Internal, "getting all boards: %v", err)
	}

	mappedBoards := utils.Map(
		allBoards,
		mappers.MapDomainBoardToProto,
	)

	return &desc.GetAllBoardsResponse{Boards: mappedBoards}, nil
}
