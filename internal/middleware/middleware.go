package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		HeaderValue := c.GetHeader("Authorization")
		if HeaderValue == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized user",
			})

			return
		}

		parts := strings.Split(HeaderValue, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized user",
			})

			return
		}

		tokenString := parts[1]

		// Step 1: Extract the token string from the Authorization header
		// Expected format: "Bearer <token>"
		// If header is missing or format is wrong → reject request

		// Step 2: jwt.Parse splits the token into three parts: header.payload.signature
		// Passes the partially parsed token (header info) to the key function

		// Step 3: Key function receives the parsed token
		// Check if the signing algorithm is HMAC (HS256)
		// If wrong algorithm → reject (prevents "alg:none" attacks)
		// If correct → return JWT_SECRET to jwt.Parse

		// Step 4: jwt.Parse takes header + payload + JWT_SECRET
		// Recalculates the signature using HMAC
		// Compares recalculated signature with the signature in the token
		// If match → token is valid and untampered
		// If no match → payload was modified, reject

		// Step 5: Check token expiration (jwt.Parse does this automatically)
		// If exp claim is in the past → reject

		// Step 6: Extract user_id from the claims (payload)
		// Set user_id in Gin context so handlers can access it
		// Call c.Next() to proceed to the actual handler
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized user",
			})

			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid claims in token",
			})

			return
		}

		userID := int(claims["user_id"].(float64))
		c.Set("userID", userID)
		c.Next()

	}
}
