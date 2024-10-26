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
	SetupRoutes(router)
	server.ListenAndServe()
}

