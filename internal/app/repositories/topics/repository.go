package topics

import desc "github.com/bogi-lyceya-44/gateway/vendor.protogen/github.com/bogi-lyceya-44/task-tracker/api/topics"

type Repository struct {
	topicClientService desc.TopicServiceClient
}

func New(
	topicClientService desc.TopicServiceClient,
) *Repository {
	return &Repository{
		topicClientService: topicClientService,
	}
}
