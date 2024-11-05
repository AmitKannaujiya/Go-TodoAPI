package cmd

import (
	_ "fmt"
	c "go-todo/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


func Execute(config *c.Config) {
	router := gin.Default()
	host, port := config.App.Host, config.App.Port

	server := http.Server{
		Addr: host + ":" +port,
		ReadTimeout: time.Duration(10 * time.Second),
		WriteTimeout: time.Duration(20 * time.Second),
		Handler: router,
	}
	// local memory route
	//SetupRoutes(router)
	// setup db routes
	SetupDBRoutes(config, router)
	server.ListenAndServe()
}

