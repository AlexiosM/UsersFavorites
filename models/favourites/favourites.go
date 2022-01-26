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

func GetFavouritesFromUser(userId int64, interFav IFav) *GetRestResponse {
	user := &User{}
	var resAssets []assets.Asset
	FavDBptr := interFav.GetFavouritesDB()
	found := false

	for u, a := range *FavDBptr {
		if u.Id == userId {
			found = true
			user = &u
			resAssets = a
			break
		}
	}
	if !found {
		return &GetRestResponse{Error: "User Not Found"}
	}

	resAssets = (*FavDBptr)[*user]
	if len(resAssets) == 0 {
		return &GetRestResponse{Error: "User has no favourites"}
	}

	return &GetRestResponse{User: *user, AssetList: resAssets}
}

func AddFavouritesToUser(userId int64, favList ListOfFavourites, interFav IFav, interAs assets.IAsDB) *PostRestResponse {
	user := &User{}
	assetList := []assets.Asset{}
	assetRespList := []assets.AssetRespData{}
	AssetDBptr := interAs.GetAssetDB()
	FavDBptr := interFav.GetFavouritesDB()

	// Get user
	found := false
	for u, _ := range *FavDBptr {
		if u.Id == userId {
			found = true
			user = &u
			break
		}
	}
	if !found {
		return &PostRestResponse{Error: "User Not Found"}
	}

	// check favourite list and remove already existing

	// add favourite list to user
	for _, fav := range favList.Favourites {
		var asset assets.Asset
		var ok bool

		if asset, ok = (*AssetDBptr)[assets.AssetId(fav)]; !ok {
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
