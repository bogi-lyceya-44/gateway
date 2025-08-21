package tasks

import (
	"context"

	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/tasks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Api) DeleteTask(
	ctx context.Context,
	request *desc.DeleteTaskRequest,
) (*desc.DeleteTaskResponse, error) {
	err := a.taskService.DeleteTask(
		ctx,
		request.GetId(),
	)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return nil, status.Error(st.Code(), st.Message())
		}

		return nil, status.Errorf(codes.Internal, "deleting task: %v", err)
	}

	return &desc.DeleteTaskResponse{}, nil
}
