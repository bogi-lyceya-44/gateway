package tasks

import (
	"context"

	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/tasks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Api) GetAllTaskDependencies(
	ctx context.Context,
	request *desc.GetAllTaskDependenciesRequest,
) (*desc.GetAllTaskDependenciesResponse, error) {
	deps, err := a.taskService.GetTaskDependencies(
		ctx,
		request.GetId(),
	)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return nil, status.Error(st.Code(), st.Message())
		}

		return nil, status.Errorf(codes.Internal, "getting task deps: %v", err)
	}

	return &desc.GetAllTaskDependenciesResponse{Ids: deps}, nil
}
