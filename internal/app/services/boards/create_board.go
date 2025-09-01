package boards

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	"github.com/pkg/errors"
)

func (s *Service) CreateBoard(
	ctx context.Context,
	board models.Board,
) (int64, error) {
	createdBoard, err := s.boardRepository.CreateBoard(ctx, board)
	if err != nil {
		return 0, errors.Wrap(err, "board repositry create board")
	}

	return createdBoard, nil
}
