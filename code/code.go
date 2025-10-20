package code

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetUserTest() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		methon := c.Request.Method
		path := c.Request.URL.Path
		
		c.Next()
		log.Printf("[%s] %s", methon, path)
	})

	r.GET("/ping/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"UserID": id,
		})
	})

	r.GET("/search", func(c *gin.Context) {
		search := c.Query("q")
		if search == "" {
			fmt.Println("Invalid: ID")
		} else {
			fmt.Println("Value :", search)
		}
		c.JSON(http.StatusOK, gin.H{"search": search})
	})

	r.GET("product/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "123" {
			c.JSON(http.StatusOK, gin.H{
				"id":   id,
				"name": "Sample",
			})
		} else {
			fmt.Println("Invalid ID")
		}

	})

	r.POST("/user", func(ctx *gin.Context) {
		var user User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error Requset"})
			return
		}
		ctx.JSON(http.StatusOK, user)
	})

	r.POST("/product", func(ctx *gin.Context) {
		var product Product
		ctx.ShouldBindJSON(&product)
		ctx.JSON(http.StatusOK, product)
	})
	r.Run(":30606")
}
