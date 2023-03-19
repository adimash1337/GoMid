package routes

import (
	"github.com/adimash1337/GoMid/controllers"
	"github.com/adimash1337/GoMid/database"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())                 //+
	incomingRoutes.POST("/users/login", controllers.Login())                   //+
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin()) //+
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())      //+
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())    //+
	//
	incomingRoutes.POST("/address/add", controllers.AddAddress())          //+
	incomingRoutes.PUT("/address/edithome", controllers.EditHomeAddress()) //+
	incomingRoutes.PUT("/address/editwork", controllers.EditWorkAddress()) //+
	incomingRoutes.GET("/address/delete", controllers.DeleteAddress())     //+
	//
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	incomingRoutes.GET("/addtocart", app.AddToCart())              //+
	incomingRoutes.GET("/removeitem", app.RemoveItem())            //+
	incomingRoutes.GET("/listcart", controllers.GetItemFromCart()) //+
	incomingRoutes.GET("/cartcheckout", app.BuyFromCart())         //+
	incomingRoutes.GET("/instantbuy", app.InstantBuy())            //+
}

//func CartRoutes(incomingRoutes *gin.Engine) {
//}
//
//func AddressRoutes(incomingRoutes *gin.Engine) {
//
//}
