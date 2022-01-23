package models

import "GWI_assingment/platform2.0-go-challenge/models/assets"

func EditDescription(description string, assetId int64) *assets.RestResponse {
	a := assets.Asset{}
	resp := a.EditAssetDescription(description, assets.AssetId(assetId))
	return resp
}
