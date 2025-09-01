package tasks

import (
	"context"

	"github.com/pkg/errors"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/tasks"
)

func (r *Repository) GetTaskDependencies(
	ctx context.Context,
	id int64,
) ([]int64, error) {
	taskToGetDepsId := []int64{id}

	response, err := r.taskServiceClient.GetAllDependencies(
		ctx,
		&desc.GetAllDependenciesRequest{
			Ids: taskToGetDepsId,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "client get task deps")
	}

	allDeps := utils.Values(
		response.GetDependenciesById(),
	)[0].GetIds()

	return allDeps, nil
}
