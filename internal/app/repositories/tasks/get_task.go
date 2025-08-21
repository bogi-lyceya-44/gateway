package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/bogi-lyceya-44/gateway/internal/app/repositories/mappers"
	desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/tasks"
	"github.com/pkg/errors"
)

func (r *Repository) GetTask(
	ctx context.Context,
	id int64,
) (models.Task, error) {
	taskToGetId := []int64{id}

	response, err := r.taskServiceClient.GetTasks(
		ctx,
		&desc.GetTasksRequest{
			Ids: taskToGetId,
		},
	)
	if err != nil {
		return models.Task{}, errors.Wrap(err, "client get task")
	}

	task := response.GetTasks()[0]

	mappedTask, err := mappers.MapProtoTaskToDomain(task)
	if err != nil {
		return models.Task{}, errors.Wrap(err, "mapping proto task to domain")
	}

	return mappedTask, nil
}
