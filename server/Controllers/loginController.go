package Controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Token struct {
	Token string
}

type User struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

func Login(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 400, "message": "Post Data Error"})
		return
	} else {
		fmt.Printf("Username is :%#v\n Password is: %#v\n", user.Username, user.Password)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "login info get successful",
		"data": &Token{
			Token: "hello",
		},
	})
}
