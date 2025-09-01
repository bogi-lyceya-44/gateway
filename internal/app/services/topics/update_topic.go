package topics

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/pkg/errors"
)

func (s *Service) UpdateTopic(
	ctx context.Context,
	topic models.UpdatedTopic,
) error {
	err := s.topicRepository.UpdateTopic(ctx, topic)
	if err != nil {
		return errors.Wrap(err, "topic repo update topic")
	}

	return nil
}
