package topics

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/pkg/errors"
)

func (s *Service) CreateTopic(
	ctx context.Context,
	topic models.Topic,
) (int64, error) {
	createdTopic, err := s.topicRepository.CreateTopic(ctx, topic)
	if err != nil {
		return 0, errors.Wrap(err, "topic repo create topic")
	}

	return createdTopic, nil
}
