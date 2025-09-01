package boards

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/pkg/errors"
)

func (s *Service) UpdateBoard(
	ctx context.Context,
	board models.UpdatedBoard,
) error {
	err := s.boardRepository.UpdateBoard(ctx, board)
	if err != nil {
		return errors.Wrap(err, "board repository update board")
	}

	return nil
}
