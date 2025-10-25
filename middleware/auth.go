package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authorization Header Missing"})
			ctx.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWTSECRET")), nil
		})
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println("Username:", claims["username"])
			fmt.Println("Expire at:", claims["exp"])
			fmt.Println("position at:", claims["position"])
			ctx.Set("positionAuth", claims["position"])
		}
		ctx.Next()
	}
}

func RequireRole(requiredRole string) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        role, exists := ctx.Get("positionAuth")
        if !exists || role != requiredRole {
            ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: insufficient permissions"})
            ctx.Abort()
            return
        }
        ctx.Next()
    }
}


func RequireRolesAllow(allowedRoles ...string) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        role, exists := ctx.Get("positionAuth")
        if !exists {
            ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: role missing"})
            ctx.Abort()
            return
        }

        userRole := role.(string)
        for _, allowed := range allowedRoles {
            if userRole == allowed {
                ctx.Next()
                return
            }
        }

        ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: insufficient permissions"})
        ctx.Abort()
    }
}