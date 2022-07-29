package api

import (
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
	errorsLib "gitlab.com/peahokinet/blockchain/travel-nft/gateway/errors"
	"gitlab.com/peahokinet/blockchain/travel-nft/gateway/helpers"
	"gitlab.com/peahokinet/blockchain/travel-nft/gateway/serializers"
	pb "gitlab.com/peahokinet/blockchain/travel-nft/protobuf/gen/go"
)

const (
	userIDKey         = "id"
	userValidTokenKey = "validToken"
)

// this const is used to force the user re-login to get newest token
var VALID_TOKEN string = "2021-11-15"

// AuthMiddleware : ...
func AuthMiddleware(key string, authenticator func(c *gin.Context) (*pb.UserInfo, error)) *jwt.GinJWTMiddleware {
	mw, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:         []byte(key),
		Timeout:     1000 * time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: userIDKey,
		Realm:       "",
		TokenLookup: "header:Authorization,query:token",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*pb.UserInfo); ok {
				return jwt.MapClaims{
					userIDKey:         v.Id,
					userValidTokenKey: VALID_TOKEN,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			user, err := authenticator(c)

			if err != nil {
				return nil, err
			}
			return user, nil
		},
		HTTPStatusMessageFunc: func(err error, c *gin.Context) string {
			return err.Error()
		},
		Unauthorized: func(c *gin.Context, _ int, message string) {
			helpers.ErrorResponse(c, http.StatusUnauthorized, errorsLib.ErrorWithMessage(errorsLib.ErrInvalidCredentials, message))
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, serializers.Resp{
				Result: serializers.UserLoginResp{
					Token:   token,
					Expired: expire.Unix(),
				},
				Error: nil,
			})
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, serializers.Resp{
				Result: serializers.UserLoginResp{
					Token:   token,
					Expired: expire.Unix(),
				},
				Error: nil,
			})
		},
	})
	return mw
}
