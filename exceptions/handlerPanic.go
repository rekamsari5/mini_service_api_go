package exceptions

import (
	"net/http"
	"service-mini-restapi/helper"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandlerRecovery(c *gin.Context, err any) {

	if errorValidation(c, err) {
		return
	}

	internalServerError(c, err)
}

func errorValidation(c *gin.Context, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		response := helper.APIResponseError("Error validate", "08", exception.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return true
	}
	return false
}

func internalServerError(c *gin.Context, err interface{}) {

	response := helper.APIResponseError("Internal Server Error", "99", "Something went wrong")
	c.JSON(http.StatusInternalServerError, response)
}
