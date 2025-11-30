package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kevinsiek889/enterprise-integrator-app-be/internal/handlers"

	_ "github.com/kevinsiek889/enterprise-integrator-app-be/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

func main() {
	r := gin.Default()

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Endpoints
	r.GET("/ping", handlers.PingHandler)

	r.Run(":8080")
}