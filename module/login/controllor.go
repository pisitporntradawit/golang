package login

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Controllor struct{
	service *Service
}

func NewControllor(service *Service) *Controllor{
	return &Controllor{
		service: service,
	}
}

func(ctrl *Controllor) GetLogin(c *gin.Context) {
	var req User
	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"errors" : "Invalid Request"})
		return 
	}
	token, err := ctrl.service.GetLogin(req.Name,req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return
	}

	c.JSON(http.StatusOK,gin.H{"token" : token})
}