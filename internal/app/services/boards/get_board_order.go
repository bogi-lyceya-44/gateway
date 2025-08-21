package boards

import (
	"context"

	"github.com/pkg/errors"
)

func (s *Service) GetBoardOrder(
	ctx context.Context,
) (map[int64]int32, error) {
	order, err := s.boardRepository.GetBoardOrder(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "board repository get board order")
	}

	return order, nil
}
