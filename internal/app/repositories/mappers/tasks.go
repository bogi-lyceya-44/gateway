package mappers

import (
	"fmt"

	"github.com/bogi-lyceya-44/common/pkg/bimap"
	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/tasks"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var mapPriority = bimap.NewFromMap(map[desc.Priority]models.Priority{
	desc.Priority_PRIORITY_LOW:      models.PriorityLow,
	desc.Priority_PRIORITY_MEDIUM:   models.PriorityMedium,
	desc.Priority_PRIORITY_HIGH:     models.PriorityHigh,
	desc.Priority_PRIORITY_CRITICAL: models.PriorityCritical,
})

func MapTaskPriorityToDomain(priority desc.Priority) (models.Priority, error) {
	value, ok := mapPriority.Get(priority)
	if !ok {
		return models.PriorityUndefined, fmt.Errorf("incorrect priority value: %v", priority)
	}

	return value, nil
}

func MapTaskPriorityToProto(priority models.Priority) (desc.Priority, error) {
	value, ok := mapPriority.GetInverse(priority)
	if !ok {
		return desc.Priority_PRIORITY_UNSPECIFIED, fmt.Errorf("incorrect priority value: %v", priority)
	}

	return value, nil
}

func MapProtoTaskToDomain(
	task *desc.Task,
) (models.Task, error) {
	mappedPriority, err := MapTaskPriorityToDomain(task.GetPriority())
	if err != nil {
		return models.Task{}, errors.Wrap(err, "mapping priority")
	}

	return models.Task{
		ID:           task.GetId(),
		Name:         task.GetName(),
		Description:  task.GetDescription(),
		Dependencies: task.GetDeps(),
		Priority:     mappedPriority,
		StartTime:    task.GetStartTime().AsTime(),
		FinishTime:   task.GetFinishTime().AsTime(),
		CreatedAt:    task.GetCreatedAt().AsTime(),
		UpdatedAt:    task.GetUpdatedAt().AsTime(),
	}, nil
}

func MapDomainTaskToCreateTaskPrototype(
	task models.Task,
) (*desc.CreateTasksRequest_TaskPrototype, error) {
	mappedPriority, err := MapTaskPriorityToProto(task.Priority)
	if err != nil {
		return nil, errors.Wrap(err, "mapping priority")
	}

	return &desc.CreateTasksRequest_TaskPrototype{
		Name:        task.Name,
		Description: task.Description,
		Deps:        task.Dependencies,
		Priority:    mappedPriority,
		StartTime:   timestamppb.New(task.StartTime),
		FinishTime:  timestamppb.New(task.FinishTime),
	}, nil
}

func MapDomainUpdatedTaskToUpdateTaskPrototype(
	task models.UpdatedTask,
) (*desc.UpdateTasksRequest_TaskPrototype, error) {
	mapped := &desc.UpdateTasksRequest_TaskPrototype{
		Id:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		Deps:        task.Dependencies,
	}

	if task.Priority != nil {
		mappedPriority, err := MapTaskPriorityToProto(*task.Priority)
		if err != nil {
			return nil, errors.Wrap(err, "mapping prioirty")
		}

		mapped.Priority = &mappedPriority
	}

	if task.StartTime != nil {
		mapped.StartTime = timestamppb.New(*task.StartTime)
	}

	if task.FinishTime != nil {
		mapped.FinishTime = timestamppb.New(*task.FinishTime)
	}

	return mapped, nil
}
