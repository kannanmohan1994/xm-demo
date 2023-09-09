package middleware

import (
	"xm/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lestrrat-go/jwx/jwk"
)

type Middleware interface {
	Authorize() gin.HandlerFunc
	CORS() gin.HandlerFunc
	CreateJWTTokenWithClaims(claims jwt.Claims, key jwk.Key, kid string) (string, error)
	GetPublicKey(token *jwt.Token) (interface{}, error)
	GetKeyByKID(KID string) (jwk.Key, error)
	Trace() gin.HandlerFunc
}

type middleware struct {
	config config.Config
}

func InitMiddleware(cfg config.Config) Middleware {
	initJWKSet(cfg)

	return &middleware{
		config: cfg,
	}
}
