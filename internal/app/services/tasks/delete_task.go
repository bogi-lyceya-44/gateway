package tasks

import (
	"context"

	"github.com/pkg/errors"
)

func (s *Service) DeleteTask(
	ctx context.Context,
	id int64,
) error {
	err := s.taskRepository.DeleteTask(ctx, id)
	if err != nil {
		return errors.Wrap(err, "task repo delete task")
	}

	return nil
}
