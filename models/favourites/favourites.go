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

func AddFavouritesToUser(userId int64, favList ListOfFavourites) *RestResponse {
	user := users.UsersDB[userId]
	fmt.Println(user)
	assetList := []assets.Asset{}
	assetRespList := []assets.AssetRespData{}

	for _, fav := range favList.Favourites {
		var asset assets.Asset
		var ok bool

		if asset, ok = assets.AssetsDB[assets.AssetId(fav)]; !ok {
			return &RestResponse{Error: errors.New("Wrong favorite id")}
		}
		assetList = append(assetList, asset)
		assetRespList = append(assetRespList, assets.AssetRespData{
			FavId: fav, AssetType: asset.AssetType})
	}

	fmt.Printf("\n%+v\n", assetList)
	Favorites[user] = assetList

	return &RestResponse{User: user, AssetList: assetRespList}
}

type RestResponse struct {
	User      users.User
	AssetList []assets.AssetRespData
	Error     error
}
