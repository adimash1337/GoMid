package routes

import (
	"github.com/adimash1337/GoMid/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	//user dimuha
	incomingRoutes.POST("/users/signup", controllers.SignUp())                 //+
	incomingRoutes.POST("/users/login", controllers.Login())                   //+
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin()) //+
	// yera
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())   //+
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery()) //+
	incomingRoutes.GET("/filter_price_asced", controllers.FilterPriceAsced())
	incomingRoutes.GET("/filter_price_desced", controllers.FilterPriceDesced())
	incomingRoutes.GET("/filter_rating_asced", controllers.FilterRatingAsced())
	incomingRoutes.GET("/filter_rating_desced", controllers.FilterRatingDesced())
	incomingRoutes.PUT("/comment/:uid/:pid", controllers.SetComment())
	//cart dimash
	incomingRoutes.GET("/addtocart", controllers.AddToCart())
	incomingRoutes.GET("/removeitem", controllers.RemoveItem())
	incomingRoutes.GET("/listcart", controllers.GetItemFromCart())
}
