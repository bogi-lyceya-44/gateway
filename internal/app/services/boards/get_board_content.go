package boards

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/pkg/errors"
)

func (s *Service) GetBoardContent(
	ctx context.Context,
	id int64,
) ([]models.TopicWithTasks, error) {
	content, err := s.boardRepository.GetBoardContent(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "board repository get board content")
	}

	return content, nil
}
