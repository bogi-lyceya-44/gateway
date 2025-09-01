package mappers

import (
	"github.com/bogi-lyceya-44/gateway/internal/app/models"
	desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/boards"
)

func MapProtoBoardToDomain(
	board *desc.Board,
) models.Board {
	return models.Board{
		ID:        board.GetId(),
		Name:      board.GetName(),
		TopicIDs:  board.GetTopicIds(),
		CreatedAt: board.GetCreatedAt().AsTime(),
		UpdatedAt: board.GetUpdatedAt().AsTime(),
	}
}

func MapDomainBoardToCreateBoardPrototype(
	board models.Board,
) *desc.CreateBoardsRequest_BoardPrototype {
	return &desc.CreateBoardsRequest_BoardPrototype{
		Name:     board.Name,
		TopicIds: board.TopicIDs,
	}
}

func MapDomainUpdatedBoardToUpdateBoardPrototype(
	board models.UpdatedBoard,
) *desc.UpdateBoardsRequest_BoardPrototype {
	return &desc.UpdateBoardsRequest_BoardPrototype{
		Id:       board.ID,
		Name:     board.Name,
		TopicIds: board.TopicIDs,
	}
}
