package api

import (
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/peahokinet/blockchain/travel-nft/gateway/config"
	errorsLib "gitlab.com/peahokinet/blockchain/travel-nft/gateway/errors"
	"gitlab.com/peahokinet/blockchain/travel-nft/gateway/services"
	pb "gitlab.com/peahokinet/blockchain/travel-nft/protobuf/gen/go"
)

// Server definition
type Server struct {
	g       *gin.Engine
	userSvc *services.UserSvc
	config  *config.Config
}

// NewServer represents for a server struct
func NewServer(
	g *gin.Engine,
	userSvc *services.UserSvc,
	config *config.Config) *Server {
	return &Server{
		g:       g,
		userSvc: userSvc,
		config:  config,
	}
}

func (s *Server) pagingFromContext(c *gin.Context) *pb.PaginationReq {

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}

	return &pb.PaginationReq{
		PageNumber: int32(page),
		PageLimit:  int32(limit),
	}
}

func (s *Server) userFromContext(c *gin.Context) (*pb.UserInfo, error) {
	userIDVal, ok := c.Get(userIDKey)
	if !ok {
		return nil, errors.New("failed to get userIDKey from context")
	}

	claims := jwt.ExtractClaims(c)
	valid, ok := claims[userValidTokenKey]
	if !ok {
		return nil, errors.New("not found user valid token key")
	}

	if valid.(string) != VALID_TOKEN {
		return nil, errors.New("invalid token")
	}

	userID := userIDVal.(string)
	user, err := s.userSvc.FindUser(userID)
	if err != nil {
		return nil, errorsLib.ErrorWithMessage(err, errors.Wrap(err, "s.userSvc.FindUser").Error())
	}

	return user, nil
}
