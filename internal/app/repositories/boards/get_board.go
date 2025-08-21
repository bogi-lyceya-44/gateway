package boards

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/bogi-lyceya-44/gateway/internal/app/repositories/mappers"
	desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/boards"
	"github.com/pkg/errors"
)

func (b *Repository) GetBoard(
	ctx context.Context,
	id int64,
) (models.Board, error) {
	boardToGet := []int64{id}

	response, err := b.boardServiceClient.GetBoards(
		ctx,
		&desc.GetBoardsRequest{
			Ids: boardToGet,
		},
	)
	if err != nil {
		return models.Board{}, errors.Wrap(err, "client get board")
	}

	board := mappers.MapProtoBoardToDomain(
		response.GetBoards()[0],
	)

	return board, nil
}
