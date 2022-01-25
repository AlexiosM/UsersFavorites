package favourites

import (
	"GWI_assingment/platform2.0-go-challenge/models/assets"
	"GWI_assingment/platform2.0-go-challenge/models/users"
	"fmt"
)

var Favorites = make(map[users.User][]assets.Asset)

type IFav interface {
	GetFavouritesDB()
}
type Fav struct {
	FavDB *map[users.User][]assets.Asset
}

func (f *Fav) GetFavouritesDB() {
	f.FavDB = &Favorites
}

// Contains the IDs of the favourite Assets
type ListOfFavourites struct {
	Favourites []int64 `json:"favourites"`
}

type PostRestResponse struct {
	User      users.User
	AssetList []assets.AssetRespData
	Error     string
}

type GetRestResponse struct {
	User      users.User
	AssetList []assets.Asset
	Error     string
}

func GetFavouritesFromUser(userId int64) *GetRestResponse {
	user := users.User{Id: userId}
	var resAssets []assets.Asset
	var ok bool

	if !user.GetUserById() {
		return &GetRestResponse{Error: "User not found"}
	}

	F := Fav{}
	F.GetFavouritesDB()
	if resAssets, ok = (*F.FavDB)[user]; !ok {
		return &GetRestResponse{Error: "No favorites for this user"}
	}

	return &GetRestResponse{User: user, AssetList: resAssets}
}

func AddFavouritesToUser(userId int64, favList ListOfFavourites) *PostRestResponse {
	assetList := []assets.Asset{}
	assetRespList := []assets.AssetRespData{}

	u := users.User{Id: userId}
	if !u.CheckIdInSlice() {
		return &PostRestResponse{Error: "User not DB"}
	}
	U := users.UsDB{}
	U.GetUserDB()

	user := (*U.Us)[userId]

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
	F := Fav{}
	F.GetFavouritesDB()
	(*F.FavDB)[user] = assetList
	fmt.Printf("\nUser %d with Asset List:", user.Id)
	for _, i := range assetList {
		fmt.Printf("\n%+v", i)
	}

	return &PostRestResponse{User: user, AssetList: assetRespList}
}
