package gincontext

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserId(c *gin.Context) (uint, error) {
	userId, exists := c.Get("x-user-id")
	if !exists {
		return 0, fmt.Errorf("User id missing")
	}

	userIdstring := userId.(string)
	userIdUint, err := strconv.ParseUint(userIdstring, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(userIdUint), nil
}
