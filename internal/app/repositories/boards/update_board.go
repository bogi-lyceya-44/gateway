package boards

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/bogi-lyceya-44/gateway/internal/app/repositories/mappers"
	desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/boards"
	"github.com/pkg/errors"
)

func (b *Repository) UpdateBoard(
	ctx context.Context,
	board models.UpdatedBoard,
) error {
	boardToUpdate := []*desc.UpdateBoardsRequest_BoardPrototype{
		mappers.MapDomainUpdatedBoardToUpdateBoardPrototype(board),
	}
	_, err := b.boardServiceClient.UpdateBoards(
		ctx,
		&desc.UpdateBoardsRequest{
			BoardsToUpdate: boardToUpdate,
		},
	)
	if err != nil {
		return errors.Wrap(err, "client update board")
	}

	return nil
}
