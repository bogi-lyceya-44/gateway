package topics

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/bogi-lyceya-44/gateway/internal/app/repositories/mappers"
	desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/topics"
	"github.com/pkg/errors"
)

func (r *Repository) CreateTopic(
	ctx context.Context,
	topic models.Topic,
) (int64, error) {
	topicToCreate := []*desc.CreateTopicsRequest_TopicPrototype{
		mappers.MapDomainTopicToCreateTopicPrototype(topic),
	}

	response, err := r.topicClientService.CreateTopics(
		ctx,
		&desc.CreateTopicsRequest{
			TopicsToCreate: topicToCreate,
		},
	)
	if err != nil {
		return 0, errors.Wrap(err, "client create topic")
	}

	createdTopicId := response.GetIds()[0]

	return createdTopicId, nil
}
