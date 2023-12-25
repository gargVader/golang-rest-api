package main

import (
	"bookstore/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	routes.RegisterBookStoreRoutes(router)

	router.Run("localhost:8080")
}
