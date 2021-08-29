package util

import (
	"errors"
	model "iris-todos/models"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/spf13/viper"
)

type Claims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func GenerateJWT(user model.User) (string, error) {
	signer := jwt.NewSigner(jwt.HS256, viper.Get("JWT_SECRET"), 60*time.Minute)

	claims := Claims{
		ID:       user.ID,
		Username: user.Username,
		Role:     "Owner",
	}

	token, err := signer.Sign(claims)
	if err != nil {
		return "", errors.New("generate JWT error")
	}

	return string(token), nil
}

func VerifyTokenHandler() context.Handler {
	verifier := jwt.NewVerifier(jwt.HS256, viper.Get("JWT_SECRET"))
	return verifier.Verify(func() interface{} {
		return new(Claims)
	})
}

func GetClaims(ctx iris.Context) *Claims {
	return jwt.Get(ctx).(*Claims)
}
