package utils

import (
	"GWI_assingment/platform2.0-go-challenge/models/assets"
	"GWI_assingment/platform2.0-go-challenge/models/favourites"
	"errors"
)

type RestResponse struct {
	AssetType   string
	Description string
	Error       error
}

func EditAssetDescription(newDescription string, id assets.AssetId) *RestResponse {
	var a assets.Asset

	// Change Assets DB description for id
	if _, ok := assets.AssetsDB[id]; ok {
		a = assets.AssetsDB[id]
		a.Description = newDescription
		assets.AssetsDB[id] = a
	} else {
		return &RestResponse{"", "", errors.New("Invalid Id")}
	}

	// Change Favourites DB description for all users that have the id
	for user, assetList := range favourites.Favorites {
		alist := []assets.Asset{}
		for _, asset := range assetList {
			if asset.AssetID == id {
				alist = favourites.Favorites[user]
			}
		}
		for index, asset := range alist {
			a := &assets.Asset{}
			if asset.AssetID == id {
				*a = asset
				a.Description = newDescription
				alist[index] = *a
				favourites.Favorites[user] = alist
			}
		}
	}

	return &RestResponse{a.AssetType, a.Description, nil}
}
