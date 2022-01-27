package favourites

import (
	"GWI_assingment/platform2.0-go-challenge/logger"
	"GWI_assingment/platform2.0-go-challenge/models/assets"
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadUsers(filepath string) {

	Users := []User{}

	jsonFile, err := os.Open(filepath)
	if err != nil {
		logger.Log.Error(err.Error())
		os.Exit(1)
	}
	logger.Log.Info("Successfully Opened users.json")
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		logger.Log.Error(err.Error())
	}
	if err := json.Unmarshal(byteValue, &Users); err != nil {
		logger.Log.Error(err.Error())
	}

	F := Fav{} // I won't test LoadUsers
	FavDBptr := F.GetFavouritesDB()

	for _, user := range Users {
		(*FavDBptr)[user] = []assets.Asset{}
	}

	// fmt.Printf("Found %d users\n", len(Users))
	// for i := 0; i < len(Users); i++ {
	// 	fmt.Println("id: " + strconv.FormatInt(Users[i].Id, 10))
	// 	fmt.Println("User firstname: " + Users[i].FirstName)
	// 	fmt.Println("User lastname: " + (Users[i].LastName))
	// 	fmt.Println("User email: " + Users[i].Email)
	// 	fmt.Println("-------------------------------")
	// }
	// fmt.Println(Favorites)
	// fmt.Println("-------------------------------")
	//
	// fmt.Println(*FavDBptr)
}
