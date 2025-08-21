package bootstrap

import (
	"log"

	"github.com/bogi-lyceya-44/common/pkg/closer"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitClientConnection(url string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, errors.Wrap(err, "init grpc client connection")
	}

	if err := closer.AddCallback(
		CloserConnGroup,
		func() error {
			log.Println("close grpc client connection")
			return conn.Close()
		},
	); err != nil {
		return nil, errors.Wrap(err, "grpc client callback")
	}

	log.Printf("grpc client connected to %s", url)

	return conn, nil
}
