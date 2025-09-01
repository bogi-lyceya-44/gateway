package mappers

import (
	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	desc "github.com/bogi-lyceya-44/gateway/internal/pb/api/gateway/boards"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapCreateBoardPrototypeToDomain(
	board *desc.CreateBoardRequest_BoardPrototype,
) models.Board {
	return models.Board{
		Name:     board.GetName(),
		TopicIDs: board.GetTopicIds(),
	}
}

func MapUpdatedBoardPrototypeToDomain(
	id int64,
	board *desc.UpdateBoardRequest_BoardPrototype,
) models.UpdatedBoard {
	return models.UpdatedBoard{
		ID:       id,
		Name:     board.Name,
		TopicIDs: board.TopicIds,
	}
}

func MapDomainBoardToProto(
	board models.Board,
) *desc.Board {
	return &desc.Board{
		Id:        board.ID,
		Name:      board.Name,
		TopicIds:  board.TopicIDs,
		CreatedAt: timestamppb.New(board.CreatedAt),
		UpdatedAt: timestamppb.New(board.UpdatedAt),
	}
}
