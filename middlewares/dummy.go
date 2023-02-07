package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Dummy(c *gin.Context) {
	fmt.Println("Im a dummy middleware " + time.Now().String())
	// Pass on to the next-in-chain
	c.Next()
}
