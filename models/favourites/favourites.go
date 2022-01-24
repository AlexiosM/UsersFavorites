package favourites

import (
	"GWI_assingment/platform2.0-go-challenge/models/assets"
	"GWI_assingment/platform2.0-go-challenge/models/users"
	"errors"
	"fmt"
)

var Favorites = make(map[users.User][]assets.Asset)

// Contains the IDs of the favourite Assets
type ListOfFavourites struct {
	Favourites []int64 `json:"favourites"`
}

type PostRestResponse struct {
	User      users.User
	AssetList []assets.AssetRespData
	Error     error
}

type GetRestResponse struct {
	User      users.User
	AssetList []assets.Asset
	Error     error
}

func GetFavouritesFromUser(userId int64) *GetRestResponse {
	user := users.User{Id: userId}
	var resAssets []assets.Asset
	var ok bool

	if !user.GetUserById() {
		return &GetRestResponse{Error: errors.New("User not found")}
	}

	if resAssets, ok = Favorites[user]; !ok {
		return &GetRestResponse{Error: errors.New("No favorites for this user")}
	}

	return &GetRestResponse{User: user, AssetList: resAssets}
}

func AddFavouritesToUser(userId int64, favList ListOfFavourites) *PostRestResponse {
	user := users.UsersDB[userId]
	assetList := []assets.Asset{}
	assetRespList := []assets.AssetRespData{}

	for _, fav := range favList.Favourites {
		var asset assets.Asset
		var ok bool

		if asset, ok = assets.AssetsDB[assets.AssetId(fav)]; !ok {
			return &PostRestResponse{Error: errors.New("Asset not inside AssetDB")}
		}
		assetList = append(assetList, asset)
		assetRespList = append(assetRespList, assets.AssetRespData{
			FavId: fav, AssetType: asset.AssetType})
	}
	Favorites[user] = assetList
	fmt.Printf("User %d with Asset List:\n%+v\n", user.Id, Favorites[user])

	return &PostRestResponse{User: user, AssetList: assetRespList}
}
