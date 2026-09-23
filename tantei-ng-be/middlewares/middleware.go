package middlewares

import (
	"strings"
	"tantei-ng/jsonwebtoken"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawHeader := c.GetHeader("Authorization")

		authHeader := strings.TrimPrefix(rawHeader, "Bearer ")

		token, err := jwt.Parse(authHeader, func(t *jwt.Token) (any, error) {
			return &jsonwebtoken.PriKey.PublicKey, nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodES256.Alg()}))

		if err != nil {
			panic(err)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("claims", claims)

			c.Next()
		}
	}

}
