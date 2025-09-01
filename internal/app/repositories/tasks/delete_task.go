package tasks

import (
	"context"

	"github.com/pkg/errors"

	desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/tasks"
)

func (r *Repository) DeleteTask(
	ctx context.Context,
	id int64,
) error {
	taskToDeleteId := []int64{id}

	_, err := r.taskServiceClient.DeleteTasks(
		ctx,
		&desc.DeleteTasksRequest{
			Ids: taskToDeleteId,
		},
	)
	if err != nil {
		return errors.Wrap(err, "client delete tasks")
	}

	return nil
}
