package route

import (
	"api/module/login"
	"api/module/user"

	"github.com/gin-gonic/gin"
)

func RouteAPI(controllor *user.Controllor,loginControllor *login.Controllor ) *gin.Engine{
	r := gin.Default()
	userGroup := r.Group("/users")
	userGroup.GET("", controllor.GetUser)
	userGroup.GET("/:id", controllor.GetUserByID)
	userGroup.POST("/", controllor.InsertUser)
	userGroup.DELETE("/:id", controllor.DeleteUser)

	r.POST("/login", loginControllor.GetLogin)
	return r
}
