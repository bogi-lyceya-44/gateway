package boards

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/pkg/errors"
)

func (s *Service) GetBoard(
	ctx context.Context,
	id int64,
) (models.Board, error) {
	board, err := s.boardRepository.GetBoard(ctx, id)
	if err != nil {
		return models.Board{}, errors.Wrap(err, "board repository get board")
	}

	return board, nil
}
