package utils

import (
	"GWI_assingment/platform2.0-go-challenge/logger"
	"GWI_assingment/platform2.0-go-challenge/models/assets"
	"GWI_assingment/platform2.0-go-challenge/models/favourites"
	"strconv"
)

type RestResponse struct {
	AssetType   string
	Description string
	Error       string
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
		logger.Log.Error("Invalid ID")
		return &RestResponse{"", "", "Invalid Id"}
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
	logger.Log.Info("AssetId " + strconv.FormatInt((int64)(a.AssetID), 10) + "(" + a.AssetType + ") took a new description: " + a.Description)

	return &RestResponse{a.AssetType, a.Description, ""}
}
