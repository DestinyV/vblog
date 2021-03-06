package Controllers

import (
	"net/http"

	"server/Services"

	"github.com/gin-gonic/gin"
)

type Return struct {
	Name    string
	Content string
}

func TestInsert(c *gin.Context) {
	var testService Services.Test

	err := c.ShouldBindJSON(&testService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := testService.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Insert() error!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "success",
		"data":    id,
	})
}

func TestGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "get",
		"data": &Return{
			Name:    "hello",
			Content: "world",
		},
	})
}
