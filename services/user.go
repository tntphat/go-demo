package services

import (
	"context"

	"gitlab.com/peahokinet/blockchain/travel-nft/gateway/config"
	"gitlab.com/peahokinet/blockchain/travel-nft/gateway/serializers"
	pb "gitlab.com/peahokinet/blockchain/travel-nft/protobuf/gen/go"
)

// UserSvc : struct
type UserSvc struct {
	conf *config.Config
}

// NewUserService w/ config
func NewUserService(conf *config.Config) *UserSvc {
	return &UserSvc{
		conf: conf,
	}
}

// FindUser by user id
func (u *UserSvc) FindUser(id string) (*pb.UserInfo, error) {
	conn, err := dialRPCConnection(u.conf.UserServiceEndpoint)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := pb.NewUserSrvClient(conn)
	res, err := c.Read(context.Background(), &pb.ReadUserReq{})
	return res.User, err
}

// Authenticate user
func (u *UserSvc) Authenticate(req *serializers.UserLoginReq) (*pb.UserInfo, error) {
	conn, err := dialRPCConnection(u.conf.UserServiceEndpoint)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := pb.NewUserSrvClient(conn)
	res, err := c.Login(context.Background(), &pb.LoginReq{})

	return res.User, err
}
