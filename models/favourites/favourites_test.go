package favourites_test

import (
	"GWI_assingment/platform2.0-go-challenge/models/assets"
	"GWI_assingment/platform2.0-go-challenge/models/favourites"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFavouritesFromUserSuccess(t *testing.T) {

	MockF := MockFav{}
	expectedDB := MockF.GetFavouritesDB()
	expectedRestResponse := &favourites.GetRestResponse{
		User:      *MockF.GetUserById(1),
		AssetList: (*expectedDB)[*MockF.GetUserById(1)],
		Error:     "",
	}

	actualResponse := favourites.GetFavouritesFromUser(1, &MockF)

	assert.EqualValues(t, *expectedRestResponse, *actualResponse)

}

func TestGetFavouritesFromUserFailureUserNotFound(t *testing.T) {
	MockDB := MockFavUserNotFound{}
	noUser := favourites.User{}
	expectedRestResponse := &favourites.GetRestResponse{
		User:      noUser,
		AssetList: nil,
		Error:     "User Not Found",
	}

	actualResponse := favourites.GetFavouritesFromUser(1, &MockDB)

	assert.EqualValues(t, *expectedRestResponse, *actualResponse)

}

func TestGetFavouritesFromUserFailureUserHasNoFavourites(t *testing.T) {
	MockDB := MockFavUserHasNoFavourites{}
	user := favourites.User{}
	assetList := []assets.Asset{}

	expectedRestResponse := &favourites.GetRestResponse{
		User:      user,
		AssetList: assetList,
		Error:     "User has no favourites",
	}

	actualResponse := favourites.GetFavouritesFromUser(1, &MockDB)

	assert.EqualValues(t, *expectedRestResponse, *actualResponse)

}

func TestAddFavouritesToUserSuccess(t *testing.T) {

	MockFavDB := MockFav{}
	MockADB := MockAs{}
	var favList favourites.ListOfFavourites
	var AssetList []assets.AssetRespData
	favList.Favourites = append(favList.Favourites, 3, 4)
	user := favourites.User{Id: 1, FirstName: "Test", LastName: "Man", Email: "somemail@example.com"}
	asset1 := assets.AssetRespData{AssetType: "insight", FavId: 3}
	asset2 := assets.AssetRespData{AssetType: "audience", FavId: 4}
	AssetList = append(AssetList, asset1, asset2)

	expectedRestResponse := &favourites.PostRestResponse{
		User:      user,
		AssetList: AssetList,
		Error:     "",
	}

	actualResponse := favourites.AddFavouritesToUser(1, favList, &MockFavDB, &MockADB)

	assert.EqualValues(t, *expectedRestResponse, *actualResponse)

}

func TestAddFavouritesToUserFailureAssetNotInsideAssetDB(t *testing.T) {
	MockDB := MockFav{}
	MockADB := MockAsFail{}
	var favList favourites.ListOfFavourites
	favList.Favourites = append(favList.Favourites, 3, 4)

	expectedRestResponse := &favourites.PostRestResponse{
		User:      favourites.User{},
		AssetList: nil,
		Error:     "Asset not inside AssetDB",
	}
	actualResponse := favourites.AddFavouritesToUser(1, favList, &MockDB, &MockADB)

	assert.EqualValues(t, *expectedRestResponse, *actualResponse)
}

func TestAddFavouritesToUserFailureUserNotFound(t *testing.T) {
	MockDB := MockFavFail{}
	MockADB := MockAs{}
	var favList favourites.ListOfFavourites
	favList.Favourites = append(favList.Favourites, 3, 4)

	expectedRestResponse := &favourites.PostRestResponse{
		User:      favourites.User{},
		AssetList: nil,
		Error:     "User Not Found",
	}
	actualResponse := favourites.AddFavouritesToUser(1, favList, &MockDB, &MockADB)

	assert.EqualValues(t, *expectedRestResponse, *actualResponse)
}

//--------- Helper Code ---------

// GET and POST Success
var MockFavourites = CreateFavouritesDB()
var MockAssetsDB = CreateAssetList()

type MockAs struct {
}

func (*MockAs) GetAssetDB() *map[assets.AssetId]assets.Asset {
	return &MockAssetsDB
}

type MockFav struct {
}

func (*MockFav) GetFavouritesDB() *map[favourites.User][]assets.Asset {
	return &MockFavourites
}

func (*MockFav) GetUserById(id int64) *favourites.User {
	MockF := MockFav{}
	for user := range *MockF.GetFavouritesDB() {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

func CreateFavouritesDB() map[favourites.User][]assets.Asset {
	result := map[favourites.User][]assets.Asset{}
	user := favourites.User{Id: 1, FirstName: "Test", LastName: "Man", Email: "somemail@example.com"}
	insight, _ := json.Marshal(assets.Insight{Text: "some text"})
	asset1 := assets.Asset{AssetType: "insight", Description: "hello insight", AssetID: 3, Asset: insight}
	audience, _ := json.Marshal(assets.Audience{Gender: "male", BirthCountry: "Greece", AgeGroup: 0, HoursInSocial: 3, Purchases: 20.0})
	asset2 := assets.Asset{AssetType: "audience", Description: "hello audience", AssetID: 4, Asset: audience}
	assetlist := []assets.Asset{}
	assetlist = append(assetlist, asset1, asset2)
	result[user] = assetlist
	return result
}

func CreateAssetList() map[assets.AssetId]assets.Asset {
	assetsmap := map[assets.AssetId]assets.Asset{}
	insight, _ := json.Marshal(assets.Insight{Text: "some text"})
	asset1 := assets.Asset{AssetType: "insight", Description: "hello insight", AssetID: 3, Asset: insight}
	audience, _ := json.Marshal(assets.Audience{Gender: "male", BirthCountry: "Greece", AgeGroup: 0, HoursInSocial: 3, Purchases: 20.0})
	asset2 := assets.Asset{AssetType: "audience", Description: "hello audience", AssetID: 4, Asset: audience}
	assetsmap[asset1.AssetID] = asset1
	assetsmap[asset2.AssetID] = asset2
	return assetsmap
}

// GET User Not Found
type MockFavUserNotFound struct {
}

func (*MockFavUserNotFound) GetFavouritesDB() *map[favourites.User][]assets.Asset {
	emptyDB := make(map[favourites.User][]assets.Asset)
	return &emptyDB
}

// GET User has no favourite
type MockFavUserHasNoFavourites struct {
}

var emptyDB = make(map[favourites.User][]assets.Asset)

func (*MockFavUserHasNoFavourites) GetFavouritesDB() *map[favourites.User][]assets.Asset {
	user := favourites.User{Id: 1, FirstName: "Test", LastName: "Man", Email: "somemail@example.com"}
	emptyDB[user] = []assets.Asset{}
	return &emptyDB
}

// POST Failure Asset not inside AssetDB

type MockAsFail struct {
}

func (*MockAsFail) GetAssetDB() *map[assets.AssetId]assets.Asset {
	empty := map[assets.AssetId]assets.Asset{}
	return &empty
}

// POST Failure User Not Found
type MockFavFail struct {
}

func (*MockFavFail) GetFavouritesDB() *map[favourites.User][]assets.Asset {
	empty := map[favourites.User][]assets.Asset{}
	return &empty
}
