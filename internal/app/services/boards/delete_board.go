package boards

import (
	"context"

	"github.com/pkg/errors"
)

func (s *Service) DeleteBoard(
	ctx context.Context,
	id int64,
) error {
	err := s.boardRepository.DeleteBoard(ctx, id)
	if err != nil {
		return errors.Wrap(err, "board repository delete board")
	}

	return nil
}
