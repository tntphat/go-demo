package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/peahokinet/blockchain/travel-nft/gateway/serializers"
)

// DefaultWelcome : ...
func (s *Server) DefaultWelcome(c *gin.Context) {
	c.JSON(http.StatusOK, "Hoki Endpoint")
}

// Welcome : ...
func (s *Server) Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, serializers.Resp{Result: "Hoki REST API"})
}
