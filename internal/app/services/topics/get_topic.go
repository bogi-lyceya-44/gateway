package topics

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/pkg/errors"
)

func (s *Service) GetTopic(
	ctx context.Context,
	id int64,
) (models.Topic, error) {
	topic, err := s.topicRepository.GetTopic(ctx, id)
	if err != nil {
		return models.Topic{}, errors.Wrap(err, "topic repo get topic")
	}

	return topic, nil
}
