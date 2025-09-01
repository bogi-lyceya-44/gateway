package boards

import (
	"context"

	desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/boards"
	"github.com/pkg/errors"
)

func (b *Repository) ChangeBoardOrder(
	ctx context.Context,
	id int64,
	place int32,
) error {
	boardOrderToChange := []*desc.BoardOrder{
		{
			BoardId: id,
			Place:   place,
		},
	}

	_, err := b.boardServiceClient.ChangeBoardOrder(
		ctx,
		&desc.ChangeBoardOrderRequest{
			Changes: boardOrderToChange,
		},
	)
	if err != nil {
		return errors.Wrap(err, "client change board order")
	}

	return nil
}
