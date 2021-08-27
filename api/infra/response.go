package infra

import (
	"github.com/gin-gonic/gin"
)

type Response struct{
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}


// ErrorJSON : json error response function
func ErrorJSON(c *gin.Context, statusCode int, message string) {
    c.JSON(statusCode, Response{Success: false, Message: message})
}

// SuccessJSON : json error response function
func SuccessJSON(c *gin.Context, statusCode int, data interface{}) {
    c.JSON(statusCode, Response{Success: true, Data: data})
}