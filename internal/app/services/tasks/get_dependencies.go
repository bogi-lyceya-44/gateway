package tasks

import (
	"context"

	"github.com/pkg/errors"
)

func (s *Service) GetTaskDependencies(
	ctx context.Context,
	id int64,
) ([]int64, error) {
	deps, err := s.taskRepository.GetTaskDependencies(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "task repo get deps")
	}

	return deps, nil
}
