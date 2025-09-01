package boards

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/bogi-lyceya-44/gateway/internal/app/repositories/mappers"
	desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/boards"
	"github.com/pkg/errors"
)

func (b *Repository) GetBoardContent(
	ctx context.Context,
	id int64,
) ([]models.TopicWithTasks, error) {
	contentToGet := []int64{id}

	response, err := b.boardServiceClient.GetBoardContent(
		ctx,
		&desc.GetBoardContentRequest{
			Ids: contentToGet,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "client get board content")
	}

	topics := utils.Values(
		response.GetContentById(),
	)[0].GetTopics()

	mapped, err := utils.MapWithError(
		topics,
		mappers.MapProtoTopicWithTasksToDomain,
	)
	if err != nil {
		return nil, errors.Wrap(err, "mapping topic to domain")
	}

	return mapped, nil
}
