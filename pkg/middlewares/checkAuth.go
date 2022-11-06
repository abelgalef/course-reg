package middlewares

import (
	"net/http"
	"strings"

	"github.com/abelgalef/course-reg/pkg/services"
	"github.com/abelgalef/course-reg/pkg/storage/repos"
	"github.com/gin-gonic/gin"
)

func CheckAuth(jwtService services.JWTTokenService, userRepo repos.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Split(c.GetHeader("Authorization"), " ")

		if len(c.GetHeader("Authorization")) == 0 || len(token) != 2 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Provide a valid Authorization header"})
			return
		}

		user, authenticated := jwtService.ValidateToken(token[1])
		if !authenticated {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User is not authenticated"})
		} else {
			usr, err := userRepo.GetUserWithID(uint(user.(map[string]interface{})["id"].(float64)))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User information not found. Your token might have been invalidated, try logging in again."})
			}
			c.Set("user", usr)
			user = nil
			c.Next()
		}
	}
}
