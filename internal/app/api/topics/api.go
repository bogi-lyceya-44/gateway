package topics

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/topics"
)

type topicService interface {
	CreateTopic(ctx context.Context, task models.Topic) (int64, error)
	DeleteTopic(ctx context.Context, id int64) error
	GetTopic(ctx context.Context, id int64) (models.Topic, error)
	UpdateTopic(ctx context.Context, task models.UpdatedTopic) error
}

type Api struct {
	desc.UnimplementedTopicServiceServer

	topicService topicService
}

func New(
	topicService topicService,
) *Api {
	return &Api{
		topicService: topicService,
	}
}
