package boards

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/pkg/errors"
)

func (s *Service) GetAllBoards(
	ctx context.Context,
) ([]models.Board, error) {
	allBoards, err := s.boardRepository.GetAllBoards(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "board repository get all boards")
	}

	return allBoards, nil
}
