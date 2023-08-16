package routers

import (
	"database/sql"
	"errors"
	"net/http"
	"service-mini-restapi/constants"
	"service-mini-restapi/database"
	"service-mini-restapi/exceptions"
	"service-mini-restapi/helper"
	"service-mini-restapi/middleware"

	formatter "github.com/abdullahPrasetio/validation-formatter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type routes struct {
	router             *gin.Engine
	db                 *sql.DB
	validatorFormatter formatter.ValidateFormatter
	validate           *validator.Validate
}

func SetupRouter() *gin.Engine {
	db, err := database.GetConnection()
	helper.PanicIfError(err)

	r := routes{
		router:   gin.Default(),
		db:       db,
		validate: formatter.Validate,
	}
	r.router.Use(gin.CustomRecovery(exceptions.ErrorHandlerRecovery))
	r.router.Use(cors.Default())
	api := r.router.Group(constants.ServerDefaultRoute, middleware.AddDefaultHeader())
	r.addCustomerRoute(api)
	r.router.NoRoute(func(c *gin.Context) {
		c.JSON(404, helper.APIResponseError("Page Not Found", constants.ErrorPageNotFound, errors.New("Page Not Found").Error()))
	})
	r.router.GET("/healthz", checkHealtz)
	return r.router
}

func checkHealtz(c *gin.Context) {
	c.JSON(http.StatusOK, helper.APIResponseSuccess("success", "ok"))
}
