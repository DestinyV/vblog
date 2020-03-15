package Controllers

import (
	"net/http"
	"server/Models"
	"server/Services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user Services.User
	var searchUser *Models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(200, gin.H{"code": 400, "message": "请求参数错误"})
		return
	}
	searchUser, err = user.Login()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    401,
			"message": "用户名或密码错误",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "登录成功",
			"data": gin.H{
				"token": "login-token",
				"data": gin.H{
					"name": searchUser.Name,
				},
			},
		})
	}
}
