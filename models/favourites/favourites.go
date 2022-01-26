package favourites

import (
	"GWI_assingment/platform2.0-go-challenge/models/assets"
	"fmt"
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
	fmt.Println(userId)
	user := &User{}
	var resAssets []assets.Asset
	var ok bool

	FavDBptr := interFav.GetFavouritesDB()

	for u, a := range *FavDBptr {
		if u.Id == userId {
			user = &u
			resAssets = a
		}
	}

	if resAssets, ok = (*FavDBptr)[*user]; !ok {
		return &GetRestResponse{Error: "User Not Found"}
	}
	if len(resAssets) == 0 {
		return &GetRestResponse{Error: "User has no favourites"}
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
