package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/pkg/errors"
)

func (s *Service) GetTask(
	ctx context.Context,
	id int64,
) (models.Task, error) {
	task, err := s.taskRepository.GetTask(ctx, id)
	if err != nil {
		return models.Task{}, errors.Wrap(err, "task repository get task")
	}

	return task, nil
}
