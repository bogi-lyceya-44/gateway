package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/pkg/errors"
)

func (s *Service) UpdateTask(
	ctx context.Context,
	task models.UpdatedTask,
) error {
	err := s.taskRepository.UpdateTask(ctx, task)
	if err != nil {
		return errors.Wrap(err, "task repo update task")
	}

	return nil
}
