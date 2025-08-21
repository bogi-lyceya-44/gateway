package topics

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/bogi-lyceya-44/gateway/internal/app/repositories/mappers"
	desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/topics"
	"github.com/pkg/errors"
)

func (r *Repository) UpdateTopic(
	ctx context.Context,
	topic models.UpdatedTopic,
) error {
	topicToUpdate := []*desc.UpdateTopicsRequest_TopicPrototype{
		mappers.MapDomainUpdatedTopicToUpdateTopicPrototype(topic),
	}

	_, err := r.topicClientService.UpdateTopics(
		ctx,
		&desc.UpdateTopicsRequest{
			TopicsToUpdate: topicToUpdate,
		},
	)
	if err != nil {
		return errors.Wrap(err, "client update topic")
	}

	return nil
}
