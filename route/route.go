package route

import (
	"api/module/user"
	"github.com/gin-gonic/gin"
)

func RouteAPI(controllor *user.Controllor) *gin.Engine{
	r := gin.Default()
	r.GET("/users", controllor.GetUser)
	r.POST("/users", controllor.InsertUser)
	r.DELETE("/users/:id", controllor.DeleteUser)
	return r
}
