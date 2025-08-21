package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/tasks"
)

type taskService interface {
	CreateTask(ctx context.Context, task models.Task) (int64, error)
	DeleteTask(ctx context.Context, id int64) error
	GetTaskDependencies(ctx context.Context, id int64) ([]int64, error)
	GetTask(ctx context.Context, id int64) (models.Task, error)
	UpdateTask(ctx context.Context, task models.UpdatedTask) error
}

type Api struct {
	desc.UnimplementedTaskServiceServer

	taskService taskService
}

func New(
	taskService taskService,
) *Api {
	return &Api{
		taskService: taskService,
	}
}
