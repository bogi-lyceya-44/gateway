package boards

import (
	"context"

	"github.com/pkg/errors"
)

func (s *Service) ChangeBoardOrder(
	ctx context.Context,
	id int64,
	place int32,
) error {
	err := s.boardRepository.ChangeBoardOrder(ctx, id, place)
	if err != nil {
		return errors.Wrap(err, "board repository change board order")
	}

	return nil
}
