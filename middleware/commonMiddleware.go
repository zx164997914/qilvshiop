package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	fmt.Println("中间件")
}
