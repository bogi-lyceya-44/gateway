package topics

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
)

type topicRepository interface {
	CreateTopic(ctx context.Context, task models.Topic) (int64, error)
	DeleteTopic(ctx context.Context, id int64) error
	GetTopic(ctx context.Context, id int64) (models.Topic, error)
	UpdateTopic(ctx context.Context, task models.UpdatedTopic) error
}

type Service struct {
	topicRepository topicRepository
}

func New(
	topicRepository topicRepository,
) *Service {
	return &Service{
		topicRepository: topicRepository,
	}
}
