package main

import (
	"log"
	"os"

	"github.com/Golang-Ecommerce_project/controllers"
	"github.com/Golang-Ecommerce_project/database"
	"github.com/Golang-Ecommerce_project/middleware"
	"github.com/Golang-Ecommerce_project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.client, "products"), database.UserData(database.client, "Users"))

	router := gin.New()
	routers.Use(gin.logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddTocart())
	router.GET("removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.instantBuy())

	log.Fatal(router.Run(":" + port))

}
