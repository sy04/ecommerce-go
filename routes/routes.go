package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sy04/ecommerce-go/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("users/login", controllers.Login())
	incomingRoutes.POST("admin/addproduct", controllers.ProductViewewAdmin())
	incomingRoutes.POST("users/productview", controllers.SearchProduct())
	incomingRoutes.POST("users/search", controllers.SearchProductByQuery())
}
