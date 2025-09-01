package mappers

import (
	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/topics"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapDomainTopicWithTaskToProto(
	topic models.TopicWithTasks,
) (*desc.TopicWithFetchedTasks, error) {
	mappedTasks, err := utils.MapWithError(
		topic.Tasks,
		MapDomainTaskToProto,
	)
	if err != nil {
		return nil, errors.Wrap(err, "mapping priority")
	}

	return &desc.TopicWithFetchedTasks{
		Id:        topic.ID,
		Name:      topic.Name,
		Tasks:     mappedTasks,
		CreatedAt: timestamppb.New(topic.CreatedAt),
		UpdatedAt: timestamppb.New(topic.UpdatedAt),
	}, nil
}

func MapDomainTopicToProto(
	topic models.Topic,
) *desc.Topic {
	return &desc.Topic{
		Id:        topic.ID,
		Name:      topic.Name,
		TaskIds:   topic.TaskIds,
		CreatedAt: timestamppb.New(topic.CreatedAt),
		UpdatedAt: timestamppb.New(topic.UpdatedAt),
	}
}

func MapCreateTopicPrototypeToDomain(
	topic *desc.CreateTopicRequest_TopicPrototype,
) models.Topic {
	return models.Topic{
		Name:    topic.GetName(),
		TaskIds: topic.GetTaskIds(),
	}
}

func MapUpdateTopicPrototypeToDomain(
	id int64,
	topic *desc.UpdateTopicRequest_TopicPrototype,
) models.UpdatedTopic {
	return models.UpdatedTopic{
		ID:      id,
		Name:    topic.Name,
		TaskIds: topic.TaskIds,
	}
}
