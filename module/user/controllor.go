package user

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Controllor struct {
	service *Service
}

func NewControllor(service *Service) *Controllor {
	return &Controllor{
		service: service,
	}
}

func (ctrl *Controllor) GetUser(c *gin.Context) {
	users, err := ctrl.service.GetUser(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func(ctrl *Controllor) InsertUser(c *gin.Context){
	var user User
	// Bind JSON body มาเป็น struct User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เรียก Service Layer insert user
	if err := ctrl.service.InsertUser(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// คืนค่า user ที่ถูก insert (รวม UUID id)
	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"name":     user.Name,
		"position": user.Position,
	})
}

func(ctrl *Controllor) DeleteUser(c *gin.Context){
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := ctrl.service.DeleteUser(ctx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message" : err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Success"})
}