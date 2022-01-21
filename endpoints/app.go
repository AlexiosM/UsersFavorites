package endpoints

import (
	"GWI_assingment/platform2.0-go-challenge/models/assets"
	"GWI_assingment/platform2.0-go-challenge/models/users"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	users.LoadUsers("./models/users/users.json")
	assets.LoadAssets("./models/assets/assets.json")
	MapUrls()
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
