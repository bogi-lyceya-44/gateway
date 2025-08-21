package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/bogi-lyceya-44/gateway/internal/app/repositories/mappers"
	"github.com/pkg/errors"

	desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/tasks"
)

func (r *Repository) CreateTask(
	ctx context.Context,
	task models.Task,
) (int64, error) {
	mappedTask, err := mappers.MapDomainTaskToCreateTaskPrototype(task)
	if err != nil {
		return 0, errors.Wrap(err, "mapping task to task create prototype")
	}

	taskToCreate := []*desc.CreateTasksRequest_TaskPrototype{
		mappedTask,
	}

	response, err := r.taskServiceClient.CreateTasks(
		ctx,
		&desc.CreateTasksRequest{
			TasksToCreate: taskToCreate,
		},
	)
	if err != nil {
		return 0, errors.Wrap(err, "client create tasks")
	}

	createdTaskId := response.GetIds()[0]

	return createdTaskId, nil
}
