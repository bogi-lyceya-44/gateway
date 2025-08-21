package tasks

import desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/tasks"

type Repository struct {
	taskServiceClient desc.TaskServiceClient
}

func New(
	taskServiceClient desc.TaskServiceClient,
) *Repository {
	return &Repository{
		taskServiceClient: taskServiceClient,
	}
}
