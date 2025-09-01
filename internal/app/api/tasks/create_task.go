package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/api/mappers"
	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/tasks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Api) CreateTask(
	ctx context.Context,
	request *desc.CreateTaskRequest,
) (*desc.CreateTaskResponse, error) {
	mappedTask, err := mappers.MapCreateTaskPrototypeToDomain(
		request.GetTaskToCreate(),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "mapping task: %v", err)
	}

	createdTask, err := a.taskService.CreateTask(
		ctx,
		mappedTask,
	)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return nil, status.Error(st.Code(), st.Message())
		}

		return nil, status.Errorf(codes.Internal, "creating task: %v", err)
	}

	return &desc.CreateTaskResponse{Id: createdTask}, nil
}
