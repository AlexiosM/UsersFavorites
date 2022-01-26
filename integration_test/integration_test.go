package integration_test

import (
	"GWI_assingment/platform2.0-go-challenge/models/assets"
	"GWI_assingment/platform2.0-go-challenge/models/favourites"
	"GWI_assingment/platform2.0-go-challenge/models/utils"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockAsset struct {
}

func (*MockAsset) GetAssetDB() *map[assets.AssetId]assets.Asset {
	//	empty := map[assets.AssetId]assets.Asset{}
	//	return &empty
	return &MockAssetsDB
}

// POST Failure User Not Found
type MockFavourite struct {
}

func (*MockFavourite) GetFavouritesDB() *map[favourites.User][]assets.Asset {
	//	empty := map[favourites.User][]assets.Asset{}
	//	return &empty
	return &MockFavourites
}

var MockAssetsDB = CreateAssetList()
var MockFavourites = CreateFavouritesDB()

func TestEditAssetDescriptionSuccess(t *testing.T) {
	MockAssetsDB := MockAsset{}
	MockFavDB := MockFavourite{}
	var favList favourites.ListOfFavourites
	favList.Favourites = append(favList.Favourites, 3)

	// PATCH
	expectedEditResponse := &utils.RestResponse{
		AssetType:   "insight",
		Description: "Adding a new Description",
		Error:       "",
	}
	user := favourites.User{Id: 1, FirstName: "Test", LastName: "Man", Email: "somemail@example.com"}

	// POST
	var assetRespList []assets.AssetRespData
	assetRespList = append(assetRespList, assets.AssetRespData{AssetType: "insight", FavId: 3})
	expectedPostResponse := &favourites.PostRestResponse{
		User:      user,
		AssetList: assetRespList,
		Error:     ""}

	// GET
	insight, _ := json.Marshal(assets.Insight{Text: "some text"})
	asset1 := assets.Asset{AssetType: "insight", Description: "Adding a new Description", AssetID: 3, Asset: insight}
	assetlist := []assets.Asset{}
	assetlist = append(assetlist, asset1)
	expectedGetResponse := &favourites.GetRestResponse{
		User:      user,
		AssetList: assetlist,
		Error:     "",
	}

	actualEditResponse := utils.EditAssetDescription("Adding a new Description", assets.AssetId(3), &MockFavDB, &MockAssetsDB)

	actualPostResponse := favourites.AddFavouritesToUser(1, favList, &MockFavDB, &MockAssetsDB)

	actualGetResponse := favourites.GetFavouritesFromUser(1, &MockFavDB)

	assert.EqualValues(t, *expectedEditResponse, *actualEditResponse)
	assert.EqualValues(t, *expectedPostResponse, *actualPostResponse)
	assert.EqualValues(t, *expectedGetResponse, *actualGetResponse)

}

func CreateAssetList() map[assets.AssetId]assets.Asset {
	assetsmap := map[assets.AssetId]assets.Asset{}
	insight, _ := json.Marshal(assets.Insight{Text: "some text"})
	asset1 := assets.Asset{AssetType: "insight", Description: "hello insight", AssetID: 3, Asset: insight}
	assetsmap[asset1.AssetID] = asset1
	return assetsmap
}

func CreateFavouritesDB() map[favourites.User][]assets.Asset {
	result := map[favourites.User][]assets.Asset{}
	user := favourites.User{Id: 1, FirstName: "Test", LastName: "Man", Email: "somemail@example.com"}
	insight, _ := json.Marshal(assets.Insight{Text: "some text"})
	asset1 := assets.Asset{AssetType: "insight", Description: "hello insight", AssetID: 3, Asset: insight}
	assetlist := []assets.Asset{}
	assetlist = append(assetlist, asset1)
	result[user] = assetlist
	return result
}
