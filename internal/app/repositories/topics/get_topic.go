package topics

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/bogi-lyceya-44/gateway/internal/app/repositories/mappers"
	desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/topics"
	"github.com/pkg/errors"
)

func (r *Repository) GetTopic(
	ctx context.Context,
	id int64,
) (models.Topic, error) {
	topicToGetId := []int64{id}

	response, err := r.topicClientService.GetTopics(
		ctx,
		&desc.GetTopicsRequest{
			Ids: topicToGetId,
		},
	)
	if err != nil {
		return models.Topic{}, errors.Wrap(err, "client get topic")
	}

	topic := response.GetTopics()[0]

	mappedTopic := mappers.MapProtoTopicToDomain(topic)

	return mappedTopic, nil
}
