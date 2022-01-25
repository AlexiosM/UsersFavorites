package favourites

import (
	"GWI_assingment/platform2.0-go-challenge/models/assets"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var Favorites = make(map[User][]assets.Asset)

type IFav interface {
	GetFavouritesDB() *map[User][]assets.Asset
}
type Fav struct {
}

func (*Fav) GetFavouritesDB() *map[User][]assets.Asset {
	return &Favorites
}

// Contains the IDs of the favourite Assets
type ListOfFavourites struct {
	Favourites []int64 `json:"favourites"`
}
type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
type PostRestResponse struct {
	User      User
	AssetList []assets.AssetRespData
	Error     string
}

type GetRestResponse struct {
	User      User
	AssetList []assets.Asset
	Error     string
}

//func GetUserById(userId int64, interFav IFav) *User {
//	FavDBptr := interFav.GetFavouritesDB()
//
//	for user, _ := range *FavDBptr {
//		if user.Id == userId {
//			return &user
//		}
//	}
//	return nil
//}

func GetFavouritesFromUser(userId int64, interFav IFav) *GetRestResponse {
	user := &User{}
	var resAssets []assets.Asset
	var ok bool

	FavDBptr := interFav.GetFavouritesDB()

	for u, _ := range *FavDBptr {
		if u.Id == userId {
			user = &u
		}
	}

	if user == nil {
		return &GetRestResponse{Error: "User not found"}
	}

	if resAssets, ok = (*FavDBptr)[*user]; !ok {
		return &GetRestResponse{Error: "No favorites for this user"}
	}

	return &GetRestResponse{User: *user, AssetList: resAssets}
}

func AddFavouritesToUser(userId int64, favList ListOfFavourites, interFav IFav) *PostRestResponse {
	user := &User{}
	assetList := []assets.Asset{}
	assetRespList := []assets.AssetRespData{}

	FavDBptr := interFav.GetFavouritesDB()

	for u, _ := range *FavDBptr {
		if u.Id == userId {
			user = &u
		}
	}

	if user == nil {
		return &PostRestResponse{Error: "User not found"}
	}

	for _, fav := range favList.Favourites {
		var asset assets.Asset
		var ok bool

		A := assets.AsDB{}
		A.GetAssetDB()

		if asset, ok = (*A.As)[assets.AssetId(fav)]; !ok {
			return &PostRestResponse{Error: "Asset not inside AssetDB"}
		}
		assetList = append(assetList, asset)
		assetRespList = append(assetRespList, assets.AssetRespData{
			FavId: fav, AssetType: asset.AssetType})
	}
	(*FavDBptr)[*user] = assetList
	fmt.Printf("\nUser %d with Asset List:", user.Id)
	for _, i := range assetList {
		fmt.Printf("\n%+v", i)
	}

	return &PostRestResponse{User: *user, AssetList: assetRespList}
}

func LoadUsers(filepath string) {

	Users := []User{}

	jsonFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Cound not read json file")
	}
	if err := json.Unmarshal(byteValue, &Users); err != nil {
		fmt.Println("Invalid json format")
	}

	F := Fav{} // I won't test LoadUsers
	FavDBptr := F.GetFavouritesDB()

	for _, user := range Users {
		(*FavDBptr)[user] = []assets.Asset{}
	}

	fmt.Printf("Found %d users\n", len(Users))
	for i := 0; i < len(Users); i++ {
		fmt.Println("id: " + strconv.FormatInt(Users[i].Id, 10))
		fmt.Println("User firstname: " + Users[i].FirstName)
		fmt.Println("User lastname: " + (Users[i].LastName))
		fmt.Println("User email: " + Users[i].Email)
		fmt.Println("-------------------------------")
	}
}
