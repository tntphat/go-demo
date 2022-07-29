package helpers

import (
	"context"
	"math/big"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/peahokinet/blockchain/travel-nft/gateway/serializers"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
)

// ErrorResponse return response with error
func ErrorResponse(c *gin.Context, s int, err error) {
	st, ok := status.FromError(err)
	if !ok {
		c.JSON(s, serializers.Resp{Error: err.Error()})
		return
	}
	p := st.Proto()
	c.JSON(s, serializers.Resp{Error: p})
}

func GenerateTimerEncryptKeyWithSalt(secret string, duration int64) (string, error) {
	now := time.Now().Unix()

	checkpoint := ((now / duration) + 1) * duration
	// to bytes
	cByte := big.NewInt(int64(checkpoint)).Bytes()
	cByte = append(cByte, []byte(secret)...)

	hashed, err := bcrypt.GenerateFromPassword(cByte, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func NewInternalCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 100*time.Millisecond)
}

func NewGrpcCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 200*time.Millisecond)
}

func NewCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}
