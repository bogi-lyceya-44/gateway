package bootstrap

import "github.com/bogi-lyceya-44/common/pkg/closer"

const (
	CloserAppGroup           = "app"
	CloserConnGroup          = "connections"
	CloserGlobalContextGroup = "global context"
)

const (
	HighPriority = iota
	MediumPriority
	LowPriority
)

func InitCloser() {
	closer.AddGroups([]closer.Group{
		{
			Name:     CloserAppGroup,
			Priority: HighPriority,
		},
		{
			Name:     CloserConnGroup,
			Priority: MediumPriority,
		},
		{
			Name:     CloserGlobalContextGroup,
			Priority: LowPriority,
		},
	}...)
}
