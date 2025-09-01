package boards

import (
	"context"

	desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/boards"
	"github.com/pkg/errors"
)

func (b *Repository) GetBoardOrder(
	ctx context.Context,
) (map[int64]int32, error) {
	response, err := b.boardServiceClient.GetBoardOrder(
		ctx,
		&desc.GetBoardOrderRequest{},
	)
	if err != nil {
		return nil, errors.Wrap(err, "client get board order")
	}

	order := response.GetOrder()

	return order, nil
}
