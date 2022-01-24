package models

import (
	"GWI_assingment/platform2.0-go-challenge/models/assets"
	"GWI_assingment/platform2.0-go-challenge/models/utils"
)

func EditDescription(description string, assetId int64) *utils.RestResponse {
	resp := utils.EditAssetDescription(description, assets.AssetId(assetId))
	return resp
}
