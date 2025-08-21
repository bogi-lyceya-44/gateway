package boards

import (
	"context"

	desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/boards"
	"github.com/pkg/errors"
)

func (b *Repository) DeleteBoard(
	ctx context.Context,
	id int64,
) error {
	boardToDelete := []int64{id}

	_, err := b.boardServiceClient.DeleteBoards(
		ctx,
		&desc.DeleteBoardsRequest{
			Ids: boardToDelete,
		},
	)
	if err != nil {
		return errors.Wrap(err, "client delete boards")
	}

	return nil
}
