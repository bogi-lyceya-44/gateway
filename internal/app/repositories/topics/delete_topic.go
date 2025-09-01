package topics

import (
	"context"

	desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/topics"
	"github.com/pkg/errors"
)

func (r *Repository) DeleteTopic(
	ctx context.Context,
	id int64,
) error {
	topicToDeleteId := []int64{id}

	_, err := r.topicClientService.DeleteTopics(
		ctx,
		&desc.DeleteTopicsRequest{
			Ids: topicToDeleteId,
		},
	)
	if err != nil {
		return errors.Wrap(err, "client delete topic")
	}

	return nil
}
