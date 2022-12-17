package controller

import (
	"github.com/gin-gonic/gin"
)

func FilterMiddle(c *gin.Context) {

	c.Next()

}
