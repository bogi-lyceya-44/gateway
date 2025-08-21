package boards

import (
	"context"

	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/boards"
)

type boardService interface {
	ChangeBoardOrder(ctx context.Context, id int64, place int32) error
	CreateBoard(ctx context.Context, board models.Board) (int64, error)
	DeleteBoard(ctx context.Context, id int64) error
	GetAllBoards(ctx context.Context) ([]models.Board, error)
	GetBoard(ctx context.Context, id int64) (models.Board, error)
	GetBoardContent(ctx context.Context, id int64) ([]models.TopicWithTasks, error)
	GetBoardOrder(ctx context.Context) (map[int64]int32, error)
	UpdateBoard(ctx context.Context, board models.UpdatedBoard) error
}

type Api struct {
	desc.UnimplementedBoardServiceServer

	boardService boardService
}

func New(boardService boardService) *Api {
	return &Api{
		boardService: boardService,
	}
}
