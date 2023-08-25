package customers

import (
	"fmt"
	"net/http"
	"service-mini-restapi/constants"
	"service-mini-restapi/helper"
	model "service-mini-restapi/models/customers"

	formattervalidator "github.com/abdullahPrasetio/validation-formatter"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type customerController struct {
	service            model.Service
	validatorFormatter formattervalidator.ValidateFormatter
	validate           *validator.Validate
}

func NewController(service model.Service, validatorFormatter formattervalidator.ValidateFormatter, validate *validator.Validate) *customerController {
	return &customerController{
		service:            service,
		validatorFormatter: validatorFormatter,
		validate:           validate,
	}
}

func (h *customerController) GetFromApi(c *gin.Context) {
	result, err := h.service.GetFromApi(c)

	if err != nil {
		response := helper.APIResponseError("Not found", constants.ErrorPageNotFound, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponseSuccess("Success get data customer", result)
	c.JSON(http.StatusOK, response)
	return

}

func (h *customerController) CreateCustomer(c *gin.Context) {
	var input model.RequestCustomer
	err := c.ShouldBindBodyWith(&input, binding.JSON)

	if err != nil {
		errorStr := fmt.Sprintf("JSON error: %v", err)
		response := helper.APIResponseError("Validation Error", constants.ErrorValidate, errorStr)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	result, err := h.service.Create(input)
	if err != nil {
		response := helper.APIResponseError("Error to create Employee", constants.ErrorGeneral, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponseSuccess("Successfully created Employee", result)
	c.JSON(http.StatusOK, response)
	return
}

func (h *customerController) GetCustomer(c *gin.Context) {
	var search model.SearchCustomer
	err := c.ShouldBindBodyWith(&search, binding.JSON)

	if err != nil {
		errorStr := fmt.Sprintf("JSON error: %v", err)
		response := helper.APIResponseError("Validation Error", constants.ErrorValidate, errorStr)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	result, err := h.service.GetAll(search)

	if len(result) < 1 {
		response := helper.APIResponseError("Customer Is Not Found ", constants.ErrorNotFound, "Not found")
		c.JSON(http.StatusNotFound, response)
		return
	}

	if err != nil {
		response := helper.APIResponseError("Error to create Employee", constants.ErrorGeneral, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponseSuccess("Successfully created Employee", model.Formatter(result))
	c.JSON(http.StatusOK, response)
	return
}

func (h *customerController) DeleteCustomer(c *gin.Context) {
	var input model.RequestDelete

	err := c.ShouldBindBodyWith(&input, binding.JSON)

	if err != nil {
		errorStr := fmt.Sprintf("JSON error: %v", err)
		response := helper.APIResponseError("Validation Error", constants.ErrorValidate, errorStr)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	h.service.Delete(input)
	response := helper.APIResponseSuccess("Successfully delete Employee", "")
	c.JSON(http.StatusOK, response)
	return
}

func (h *customerController) UpdateCustomer(c *gin.Context) {
	var input model.RequestUpdate
	err := c.ShouldBindBodyWith(&input, binding.JSON)

	if err != nil {
		errorStr := fmt.Sprintf("JSON error: %v", err)
		response := helper.APIResponseError("Validation Error", constants.ErrorValidate, errorStr)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	result, err := h.service.Update(input)
	if err != nil {
		response := helper.APIResponseError("Error to update Employee", constants.ErrorGeneral, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponseSuccess("Successfully update Employee", result)
	c.JSON(http.StatusOK, response)
	return
}
