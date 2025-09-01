package boards

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/bogi-lyceya-44/gateway/internal/app/repositories/mappers"
	desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/boards"
	"github.com/pkg/errors"
)

func (b *Repository) GetAllBoards(
	ctx context.Context,
) ([]models.Board, error) {
	response, err := b.boardServiceClient.GetAllBoards(
		ctx,
		&desc.GetAllBoardsRequest{},
	)
	if err != nil {
		return nil, errors.Wrap(err, "client get all boards")
	}

	allBoards := utils.Map(
		response.GetBoards(),
		mappers.MapProtoBoardToDomain,
	)

	return allBoards, nil
}
