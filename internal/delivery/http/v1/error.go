package v1

import "github.com/gin-gonic/gin"

type Response struct {
	Message string `json:"message"`
}

func errorResponse(c *gin.Context, statusCode int, err error) {
	c.AbortWithStatusJSON(statusCode, Response{Message: err.Error()})
}
