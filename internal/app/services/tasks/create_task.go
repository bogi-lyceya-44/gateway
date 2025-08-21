package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/pkg/errors"
)

func (s *Service) CreateTask(
	ctx context.Context,
	task models.Task,
) (int64, error) {
	createdTask, err := s.taskRepository.CreateTask(ctx, task)
	if err != nil {
		return 0, errors.Wrap(err, "task repo create task")
	}

	return createdTask, nil
}
