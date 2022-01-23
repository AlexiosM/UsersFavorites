package assets

import (
	"errors"
	"fmt"
)

type RestResponse struct {
	AssetType   string
	Description string
	Error       error
}

func (Asset) EditAssetDescription(newDescription string, id AssetId) *RestResponse {
	var a Asset

	if _, ok := assetsDB[id]; ok {
		a = assetsDB[id]
		a.Description = newDescription
		fmt.Println(a.AssetType)
		fmt.Println(a.Description)

	} else {
		return &RestResponse{"", "", errors.New("Invalid Id")}
	}
	return &RestResponse{a.AssetType, a.Description, nil}
}
