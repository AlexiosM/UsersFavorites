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

func EditAssetDescription(newDescription string, id assets.AssetId, interFav favourites.IFav, interAs assets.IAsDB) *RestResponse {
	var a assets.Asset

	AssetDBptr := interAs.GetAssetDB()
	// Change Assets DB description for id
	if _, ok := (*AssetDBptr)[id]; ok {
		a = (*AssetDBptr)[id]
		a.Description = newDescription
		(*AssetDBptr)[id] = a
	} else {
		return &RestResponse{"", "", errors.New("Invalid Id")}
	}

	// Change Favourites DB description for all users that have the id
	FavDBptr := interFav.GetFavouritesDB()
	for user, assetList := range *FavDBptr {
		alist := []assets.Asset{}
		for _, asset := range assetList {
			if asset.AssetID == id {
				//alist = favourites.Favorites[user]
				alist = (*FavDBptr)[user]
			}
		}
		for index, asset := range alist {
			a := &assets.Asset{}
			if asset.AssetID == id {
				*a = asset
				a.Description = newDescription
				alist[index] = *a
				(*FavDBptr)[user] = alist
			}
		}
	}

	return &RestResponse{a.AssetType, a.Description, nil}
}
