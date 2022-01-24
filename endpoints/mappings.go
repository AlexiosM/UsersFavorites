package endpoints

import "GWI_assingment/platform2.0-go-challenge/controllers"

func MapUrls() {

	router.POST("/addassettofav/:user_id", controllers.AddFavourites)
	router.GET("/getuserfavourites/:user_id", controllers.GetFavourites)
	router.PATCH("/editasset/:asset_id", controllers.ChangeDescription)
}
