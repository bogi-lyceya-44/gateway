package boards

import desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/boards"

type Repository struct {
	boardServiceClient desc.BoardServiceClient
}

func New(
	boardServiceClient desc.BoardServiceClient,
) *Repository {
	return &Repository{
		boardServiceClient: boardServiceClient,
	}
}
