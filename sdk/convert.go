package sdk

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ConvertStr(str string, c *gin.Context) int {
	result, ok := strconv.Atoi(str)
	if ok != nil {
		FailOrError(c, 500, "Failed to convert", ok)
		return 0
	}
	return result
}
