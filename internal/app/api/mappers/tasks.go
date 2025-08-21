package mappers

import (
	"fmt"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/tasks"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/bogi-lyceya-44/common/pkg/bimap"
)

var mapPriority = bimap.NewFromMap(map[models.Priority]desc.Priority{
	models.PriorityUndefined: desc.Priority_PRIORITY_UNSPECIFIED,
	models.PriorityLow:       desc.Priority_PRIORITY_LOW,
	models.PriorityMedium:    desc.Priority_PRIORITY_MEDIUM,
	models.PriorityHigh:      desc.Priority_PRIORITY_HIGH,
	models.PriorityCritical:  desc.Priority_PRIORITY_CRITICAL,
})

func MapDomainTaskPriorityToProto(priority models.Priority) (desc.Priority, error) {
	value, ok := mapPriority.Get(priority)
	if !ok {
		return desc.Priority_PRIORITY_UNSPECIFIED, fmt.Errorf("undefined priority value: %v", priority)
	}

	return value, nil
}

func MapProtoTaskPriorityToDomain(priority desc.Priority) (models.Priority, error) {
	value, ok := mapPriority.GetInverse(priority)
	if !ok {
		return models.PriorityUndefined, fmt.Errorf("undefined priority value: %v", priority)
	}

	return value, nil
}

func MapDomainTaskToProto(
	task models.Task,
) (*desc.Task, error) {
	mappedPriority, err := MapDomainTaskPriorityToProto(task.Priority)
	if err != nil {
		return nil, errors.Wrap(err, "mapping priority")
	}

	return &desc.Task{
		Id:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		Deps:        task.Dependencies,
		Priority:    mappedPriority,
		StartTime:   timestamppb.New(task.StartTime),
		FinishTime:  timestamppb.New(task.FinishTime),
		CreatedAt:   timestamppb.New(task.CreatedAt),
		UpdatedAt:   timestamppb.New(task.UpdatedAt),
	}, nil
}

func MapCreateTaskPrototypeToDomain(
	task *desc.CreateTaskRequest_TaskPrototype,
) (models.Task, error) {
	mappedPriority, err := MapProtoTaskPriorityToDomain(task.Priority)
	if err != nil {
		return models.Task{}, errors.Wrap(err, "mapping priority")
	}

	return models.Task{
		Name:         task.GetName(),
		Description:  task.GetDescription(),
		Dependencies: task.GetDeps(),
		Priority:     mappedPriority,
		StartTime:    task.GetStartTime().AsTime(),
		FinishTime:   task.GetFinishTime().AsTime(),
	}, nil
}

func MapUpdateTaskPrototypeToDomain(
	id int64,
	task *desc.UpdateTaskRequest_TaskPrototype,
) (models.UpdatedTask, error) {
	mapped := models.UpdatedTask{
		ID:           id,
		Name:         task.Name,
		Description:  task.Description,
		Dependencies: task.Deps,
	}

	if task.Priority != nil {
		mappedPriority, err := MapProtoTaskPriorityToDomain(task.GetPriority())
		if err != nil {
			return models.UpdatedTask{}, errors.Wrap(err, "mapping prioirty")
		}

		mapped.Priority = &mappedPriority
	}

	if task.StartTime != nil {
		temp := task.StartTime.AsTime()
		mapped.StartTime = &temp
	}

	if task.FinishTime != nil {
		temp := task.FinishTime.AsTime()
		mapped.FinishTime = &temp
	}

	return mapped, nil
}
