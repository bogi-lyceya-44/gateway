package mappers

import (
	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/topics"
	"github.com/pkg/errors"
)

func MapProtoTopicWithTasksToDomain(
	topic *desc.TopicWithFetchedTasks,
) (models.TopicWithTasks, error) {
	mappedTasks, err := utils.MapWithError(
		topic.GetTasks(),
		MapProtoTaskToDomain,
	)
	if err != nil {
		return models.TopicWithTasks{},
			errors.Wrap(err, "mapping priority")
	}

	return models.TopicWithTasks{
		ID:        topic.GetId(),
		Name:      topic.GetName(),
		Tasks:     mappedTasks,
		CreatedAt: topic.GetCreatedAt().AsTime(),
		UpdatedAt: topic.GetUpdatedAt().AsTime(),
	}, nil
}

func MapProtoTopicToDomain(
	topic *desc.Topic,
) models.Topic {
	return models.Topic{
		ID:        topic.GetId(),
		Name:      topic.GetName(),
		TaskIds:   topic.GetTaskIds(),
		CreatedAt: topic.GetCreatedAt().AsTime(),
		UpdatedAt: topic.GetUpdatedAt().AsTime(),
	}
}

func MapDomainUpdatedTopicToUpdateTopicPrototype(
	topic models.UpdatedTopic,
) *desc.UpdateTopicsRequest_TopicPrototype {
	return &desc.UpdateTopicsRequest_TopicPrototype{
		Id:      topic.ID,
		Name:    topic.Name,
		TaskIds: topic.TaskIds,
	}
}

func MapDomainTopicToCreateTopicPrototype(
	topic models.Topic,
) *desc.CreateTopicsRequest_TopicPrototype {
	return &desc.CreateTopicsRequest_TopicPrototype{
		Name:    topic.Name,
		TaskIds: topic.TaskIds,
	}
}
