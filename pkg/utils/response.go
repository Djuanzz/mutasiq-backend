package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Status  string      `json:"status" form:"status"`
	Message string      `json:"message" form:"message" omitempty:""`
	Data    interface{} `json:"data" form:"data" omitempty:""`
}

func SuccessResponse(c *gin.Context, httpStatus int, message string, data interface{}) {
	c.JSON(httpStatus, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, httpStatus int, message string) {
	c.JSON(httpStatus, Response{
		Status:  "error",
		Message: message,
	})
}
