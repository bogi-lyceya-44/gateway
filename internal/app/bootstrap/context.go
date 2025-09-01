package bootstrap

import (
	"context"
	"log"

	"github.com/bogi-lyceya-44/common/pkg/closer"
	"github.com/pkg/errors"
)

func InitGlobalContext() (context.Context, error) {
	ctx, cancel := context.WithCancel(context.Background())

	if err := closer.AddCallback(
		CloserGlobalContextGroup,
		func() error {
			log.Println("cancel context")
			cancel()
			return nil
		},
	); err != nil {
		return nil, errors.Wrap(err, "global context callback")
	}

	return ctx, nil
}
