package topics

import (
	"context"

	"github.com/pkg/errors"
)

func (s *Service) DeleteTopic(
	ctx context.Context,
	id int64,
) error {
	err := s.topicRepository.DeleteTopic(ctx, id)
	if err != nil {
		return errors.Wrap(err, "topic repo delete topic")
	}

	return nil
}
