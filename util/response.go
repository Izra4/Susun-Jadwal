package util

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HttpSuccessResponse(c *gin.Context, code int, msg string, data interface{}) {
	switch code / 100 {
	case 2:
		c.JSON(code, map[string]interface{}{
			"status":  "success",
			"message": msg,
			"data":    data,
		})
	default:
		c.JSON(500, map[string]interface{}{
			"status":  "error",
			"message": "Response Error",
		})
	}
}

func HttpFailOrErrorResponse(c *gin.Context, httpCode int, msg string, err error) {
	switch httpCode / 100 {
	case 4: //FAIL 4xx
		c.JSON(httpCode, gin.H{
			"status":  "fail",
			"message": msg,
			"data": gin.H{
				"error": err.Error(),
			},
		})
	case 5: //ERROR 5xx
		log.Println(err)
		c.JSON(httpCode, gin.H{
			"status":  "error",
			"message": msg,
		})

	default:
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "RESPONSE ERROR",
		})
	}
}

func ErrorEmptyField(c *gin.Context) {
	c.JSON(400, gin.H{
		"status":  "fail",
		"message": "Please fill the empty field",
	})
}

func ErrorConvertStr(str string, c *gin.Context) (int, error) {
	result, ok := strconv.Atoi(str)
	if ok != nil {
		HttpFailOrErrorResponse(c, 500, "Failed to convert", ok)
		return 0, ok
	}
	return result, nil
}