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
			fmt.Println("role at:", claims["role"])
			ctx.Set("positionAuth", claims["role"])
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

		interfaceRole, ok := role.([]interface{}) // แปลงจาก interface{} ให้เป็น slice ของ interface{}
		fmt.Println(interfaceRole)

		var roles []string // สร้าง slice ของ string เพื่อเก็บ role ที่แปลงแล้ว
		for _, r := range interfaceRole { // วน loop แปลงแต่ละค่าใน slice ให้เป็น string
			if s, ok := r.(string); ok {
				roles = append(roles, s)
			}
		}
		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error: roles type invalid"})
			ctx.Abort()
			return
		}
		for _, allowed := range allowedRoles {
			for _, userRole := range roles {
				if strings.EqualFold(userRole, allowed) { // case-insensitive
					ctx.Next()
					return
				}
			}

		}

		ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: insufficient permissions"})
		ctx.Abort()
	}
}
