package main

import (
	"fmt"
	"net/http"
	"service-mini-restapi/configs"
	"service-mini-restapi/constants"
	"service-mini-restapi/routers"

	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func init() {
	config, _ := configs.LoadConfig(".")
	gin.SetMode(config.GIN_MODE)
}

func main() {

	router := routers.SetupRouter()
	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", constants.ServerPort),
		Handler: router,
	}

	fmt.Println("Server run in url : http://localhost:" + constants.ServerPort)

	err := server.ListenAndServe()
	server.ErrorLog.Panicln(err)
}
