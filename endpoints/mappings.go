package endpoints

import "GWI_assingment/platform2.0-go-challenge/controllers"

func MapUrls() {

	// router.POST("/createuser", controllers.CreateUser)
	// router.GET("/createuser:user_id", controllers.CreateUser)
	// router.GET("/addassettofav:user_id:asset_id", controllers.CreateUser)
	router.PATCH("/editasset/:asset_id", controllers.ChangeDescription)
}
