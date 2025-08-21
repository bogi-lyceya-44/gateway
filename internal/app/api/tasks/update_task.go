package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/api/mappers"
	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/tasks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Api) UpdateTask(
	ctx context.Context,
	request *desc.UpdateTaskRequest,
) (*desc.UpdateTaskResponse, error) {
	mappedTask, err := mappers.MapUpdateTaskPrototypeToDomain(
		request.GetId(),
		request.GetTaskToUpdate(),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "mapping task: %v", err)
	}

	err = a.taskService.UpdateTask(
		ctx,
		mappedTask,
	)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return nil, status.Error(st.Code(), st.Message())
		}

		return nil, status.Errorf(codes.Internal, "updating task: %v", err)
	}

	return &desc.UpdateTaskResponse{}, nil
}
