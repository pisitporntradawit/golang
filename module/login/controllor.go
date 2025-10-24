package login

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Controllor struct {
	service *Service
}

func NewControllor(service *Service) *Controllor {
	return &Controllor{
		service: service,
	}
}

func (ctrl *Controllor) GetLogin(c *gin.Context) {
	var req User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid Request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	token, err := ctrl.service.GetLogin(ctx, req.Username, req.Password)
	if err != nil {
		if err.Error() == "invalid password" || err.Error() == "user not found" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
			return
		}
	}
	fmt.Println("Login Success Username:", req.Username)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
