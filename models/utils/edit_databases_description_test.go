package utils_test

import (
	"GWI_assingment/platform2.0-go-challenge/models/assets"
	"GWI_assingment/platform2.0-go-challenge/models/favourites"
	"GWI_assingment/platform2.0-go-challenge/models/utils"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockAsSuccess struct {
}

func (*MockAsSuccess) GetAssetDB() *map[assets.AssetId]assets.Asset {
	return &MockAssetsDB
}

type MockFavFail struct {
}

func (*MockFavFail) GetFavouritesDB() *map[favourites.User][]assets.Asset {
	empty := map[favourites.User][]assets.Asset{}
	return &empty
}

func CreateAssetList() map[assets.AssetId]assets.Asset {
	assetsmap := map[assets.AssetId]assets.Asset{}
	insight, _ := json.Marshal(assets.Insight{Text: "some text"})
	asset1 := assets.Asset{AssetType: "insight", Description: "hello insight", AssetID: 3, Asset: insight}
	assetsmap[asset1.AssetID] = asset1
	return assetsmap
}

var MockAssetsDB = CreateAssetList()

func TestEditAssetDescriptionSuccess(t *testing.T) {
	MockAssetsDB := MockAsSuccess{}
	MockFavDB := MockFavFail{} // No user is needed to change asset description

	expectedRestResponse := &utils.RestResponse{
		AssetType:   "insight",
		Description: "Adding a new Description",
		Error:       "",
	}

	actualResponse := utils.EditAssetDescription("Adding a new Description", assets.AssetId(3), &MockFavDB, &MockAssetsDB)

	assert.EqualValues(t, *expectedRestResponse, *actualResponse)

}

func TestEditAssetDescriptionFailureAskForNotExistingAssetID(t *testing.T) {
	MockAssetsDB := MockAsSuccess{}
	MockFavDB := MockFavFail{}

	expectedRestResponse := &utils.RestResponse{
		AssetType:   "",
		Description: "",
		Error:       "Invalid Id",
	}

	actualResponse := utils.EditAssetDescription("Adding a new Description", assets.AssetId(1), &MockFavDB, &MockAssetsDB)

	assert.EqualValues(t, *expectedRestResponse, *actualResponse)

}
