package boards

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/bogi-lyceya-44/gateway/internal/app/repositories/mappers"
	desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/boards"
	"github.com/pkg/errors"
)

func (b *Repository) CreateBoard(
	ctx context.Context,
	board models.Board,
) (int64, error) {
	boardToCreate := []*desc.CreateBoardsRequest_BoardPrototype{
		mappers.MapDomainBoardToCreateBoardPrototype(board),
	}

	response, err := b.boardServiceClient.CreateBoards(
		ctx,
		&desc.CreateBoardsRequest{
			BoardsToCreate: boardToCreate,
		},
	)
	if err != nil {
		return 0, errors.Wrap(err, "client create boards")
	}

	createdBoardId := response.GetIds()[0]

	return createdBoardId, nil
}
