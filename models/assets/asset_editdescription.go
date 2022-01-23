package assets

import (
	"errors"
)

type RestResponse struct {
	AssetType   string
	Description string
	Error       error
}

func (Asset) EditAssetDescription(newDescription string, id AssetId) *RestResponse {
	var a Asset

	if _, ok := AssetsDB[id]; ok {
		a = AssetsDB[id]
		a.Description = newDescription
	} else {
		return &RestResponse{"", "", errors.New("Invalid Id")}
	}
	return &RestResponse{a.AssetType, a.Description, nil}
}
