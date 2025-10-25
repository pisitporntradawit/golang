package route

import (
	"api/middleware"
	"api/module/login"
	"api/module/user"
	"github.com/gin-gonic/gin"
)

func RouteAPI(controllor *user.Controllor, loginControllor *login.Controllor) *gin.Engine {
	r := gin.Default()
	r.POST("/login", loginControllor.GetLogin)
	r.Use(middleware.AuthLogin())
	userGroup := r.Group("/users")
	userGroup.GET("", middleware.RequireRolesAllow("admin", "eng"),controllor.GetUser)
	userGroup.GET("/:id", controllor.GetUserByID)
	userGroup.POST("/", controllor.InsertUser)
	userGroup.DELETE("/:id", controllor.DeleteUser)

	return r
}
