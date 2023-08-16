package routers

import (
	"service-mini-restapi/controllers/api/v1/customers"
	models "service-mini-restapi/models/customers"

	"github.com/gin-gonic/gin"
)

func (r *routes) addCustomerRoute(rg *gin.RouterGroup) {
	repository := models.NewRepository(r.db)
	service := models.NewService(repository)
	controller := customers.NewController(service, r.validatorFormatter, r.validate)

	customer := rg.Group("customer")
	customer.GET("/from-api", controller.GetFromApi)
	customer.POST("/create", controller.CreateCustomer)
	customer.POST("/getall", controller.GetCustomer)
	customer.POST("/delete", controller.DeleteCustomer)
	customer.POST("/update", controller.UpdateCustomer)

}
