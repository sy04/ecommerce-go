package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sy04/ecommerce-go/controllers"
	"github.com/sy04/ecommerce-go/database"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())
}
