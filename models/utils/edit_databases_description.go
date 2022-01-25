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

	A := assets.AsDB{}
	A.GetAssetDB()
	// Change Assets DB description for id
	if _, ok := (*A.As)[id]; ok {
		a = (*A.As)[id]
		a.Description = newDescription
		(*A.As)[id] = a
	} else {
		return &RestResponse{"", "", errors.New("Invalid Id")}
	}

	// Change Favourites DB description for all users that have the id
	F := favourites.Fav{}
	F.GetFavouritesDB()
	for user, assetList := range *F.FavDB {
		alist := []assets.Asset{}
		for _, asset := range assetList {
			if asset.AssetID == id {
				//alist = favourites.Favorites[user]
				alist = (*F.FavDB)[user]
			}
		}
		for index, asset := range alist {
			a := &assets.Asset{}
			if asset.AssetID == id {
				*a = asset
				a.Description = newDescription
				alist[index] = *a
				(*F.FavDB)[user] = alist
			}
		}
	}

	return &RestResponse{a.AssetType, a.Description, nil}
}
