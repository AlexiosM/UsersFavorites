package endpoints

import "GWI_assingment/platform2.0-go-challenge/controllers"

func MapUrls() {

	router.POST("/addassettofav/:user_id", controllers.AddFavorites)
	// router.GET("/createuser:user_id", controllers.CreateUser)
	router.PATCH("/editasset/:asset_id", controllers.ChangeDescription)
}
