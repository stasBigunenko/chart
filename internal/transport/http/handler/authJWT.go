package handler

import (
	"chart/internal/config"
	"chart/internal/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

func (h *handler) VerifyUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		notAuth := []string{"/signup", "/signin"}
		requestPath := c.Request.URL.Path

		for _, val := range notAuth {
			if val == requestPath {
				c.Next()
				return
			}
		}

		token, err := c.Cookie("chartJWT")
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		claims, ok := validateToken(token)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("invalid token")})
			c.Abort()
			return
		}

		c.Set("jwt", *claims)
		c.Next()
	}
}

func validateToken(jwtToken string) (*models.Claims, bool) {
	claims := &models.Claims{}

	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRETKEY), nil
	})
	if err != nil {
		return claims, false
	}

	if !token.Valid {
		return claims, false
	}

	return claims, true
}

func (h *handler) generateTokenStringForUser(id, name string) (string, error) {
	// Create the JWT claims, which includes the username and expiry time
	claims := models.Claims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    id,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.SECRETKEY))
	return tokenString, err
}
