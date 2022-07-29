package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errorsLib "gitlab.com/peahokinet/blockchain/travel-nft/gateway/errors"
	"gitlab.com/peahokinet/blockchain/travel-nft/gateway/helpers"
	"gitlab.com/peahokinet/blockchain/travel-nft/gateway/serializers"
	pb "gitlab.com/peahokinet/blockchain/travel-nft/protobuf/gen/go"
)

// Authenticate : [POST] /auth/verify
func (s *Server) Authenticate(c *gin.Context) (*pb.UserInfo, error) {
	var req serializers.UserLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, errorsLib.ErrorWithMessage(errorsLib.ErrInvalidArgument, err.Error())
	}

	res, err := s.userSvc.Authenticate(&req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// RequestLogin request to get access-code
func (s *Server) RequestLogin(c *gin.Context) {
	var req serializers.UserLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, errorsLib.ErrorWithMessage(errorsLib.ErrInvalidArgument, err.Error()))
		return
	}

	res, err := s.userSvc.Authenticate(&req)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: res, Error: nil})
}

// GetUserProfile : [GET] /user/profile
func (s *Server) GetUserProfile(c *gin.Context) {
	user, err := s.userFromContext(c)

	if err != nil {
		helpers.ErrorResponse(c, http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: user, Error: nil})
}
