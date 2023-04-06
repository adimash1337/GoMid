package routes

import (
	"github.com/adimash1337/GoMid/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())                 //+
	incomingRoutes.POST("/users/login", controllers.Login())                   //+
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin()) //+
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())      //+
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())    //+
	incomingRoutes.GET("/filter_price", controllers.FilterPrice())
	incomingRoutes.GET("/filter_rating", controllers.FilterRating())
}
