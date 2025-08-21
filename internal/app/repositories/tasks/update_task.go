package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/bogi-lyceya-44/gateway/internal/app/repositories/mappers"
	desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/tasks"
	"github.com/pkg/errors"
)

func (r *Repository) UpdateTask(
	ctx context.Context,
	task models.UpdatedTask,
) error {
	mappedTask, err := mappers.MapDomainUpdatedTaskToUpdateTaskPrototype(task)
	if err != nil {
		return errors.Wrap(err, "mapping updated task to update task prototype")
	}

	taskToUpdate := []*desc.UpdateTasksRequest_TaskPrototype{
		mappedTask,
	}

	_, err = r.taskServiceClient.UpdateTasks(
		ctx,
		&desc.UpdateTasksRequest{
			TasksToUpdate: taskToUpdate,
		},
	)
	if err != nil {
		return errors.Wrap(err, "client update task")
	}

	return nil
}
