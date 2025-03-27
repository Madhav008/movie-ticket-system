package main

import (
	"fmt"
	"log"
	"movieTicket/config"
	"movieTicket/repository"
	"movieTicket/routes"
	"movieTicket/services"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	// Load configuration
	cfg := config.InitConfig()
	port := cfg.Server.Port

	router := gin.Default()

	//Adding the logging middlware
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// Adding recovery middleware to handle panics and avoid server crashes
	router.Use(gin.Recovery())

	// Define a routes group for the API endpoints
	repo := repository.NewMovieTicketRepository()
	var service services.ServiceInterface
	service = services.NewMovieTicketService(repo)
	routes.SetupRoutes(router, service)

	// Start the Gin server on port 8080
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
