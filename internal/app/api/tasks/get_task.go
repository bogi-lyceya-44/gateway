package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/api/mappers"
	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/tasks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Api) GetTask(
	ctx context.Context,
	request *desc.GetTaskRequest,
) (*desc.GetTaskResponse, error) {
	task, err := a.taskService.GetTask(
		ctx,
		request.GetId(),
	)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return nil, status.Error(st.Code(), st.Message())
		}

		return nil, status.Errorf(codes.Internal, "getting task: %v", err)
	}

	mappedTask, err := mappers.MapDomainTaskToProto(task)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "mapping task: %v", err)
	}

	return &desc.GetTaskResponse{Task: mappedTask}, nil
}
