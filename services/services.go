package services

import (
	"time"

	"github.com/pkg/errors"
	errorsLib "gitlab.com/peahokinet/blockchain/travel-nft/gateway/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const TokenLifeCycle = 300

func dialRPCConnection(serviceEndpoint string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(serviceEndpoint, grpc.WithInsecure(), grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:    10 * time.Second,
		Timeout: 500 * time.Millisecond,
	}))
	if err != nil {
		return nil, errorsLib.ErrorWithMessage(errorsLib.ErrSystemError, errors.Wrap(err, "u.Dial").Error())
	}
	return conn, nil
}
