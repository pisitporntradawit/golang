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

func (ctrl *Controllor) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "ID Not Found"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := ctrl.service.GetUserByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ID": user.ID,
		"Username" : user.Username,
		"Name" : user.Name,
		"Position" : user.Position,
})
}

func (ctrl *Controllor) InsertUser(c *gin.Context) {
	var req UserProfile
	// Bind JSON body มาเป็น struct User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := User{
		Username: req.Username,
		Password: req.Password,
	}

	profile := Profile{
		Name:     req.Name,
		Position: req.Position,
	}

	// เรียก Service Layer insert user
	if err := ctrl.service.InsertUser(c.Request.Context(), &user, &profile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// คืนค่า user ที่ถูก insert (รวม UUID id)
	c.JSON(http.StatusOK, gin.H{
		"id":        user.ID,
		"Username":  user.Username,
		"password":  user.Password,
		"name":      profile.Name,
		"position":  profile.Position,
		"profileID": profile.ID,
	})
}

func (ctrl *Controllor) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := ctrl.service.DeleteUser(ctx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
