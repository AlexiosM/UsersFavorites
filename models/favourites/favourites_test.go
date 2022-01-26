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

	expectedRestResponse := &favourites.GetRestResponse{
		User:      user,
		AssetList: nil,
		Error:     "User has no favourites",
	}

	actualResponse := favourites.GetFavouritesFromUser(1, &MockDB)

	assert.EqualValues(t, *expectedRestResponse, *actualResponse)

}

func TestAddFavouritesToUserSuccess(t *testing.T) {

}
func TestAddFavouritesToUserFailure(t *testing.T) {

}

// Helper Code

// GET Success
var MockFavourites = CreateFavouritesDB()

type MockFav struct {
}

func (*MockFav) GetFavouritesDB() *map[favourites.User][]assets.Asset {
	return &MockFavourites
}

func (*MockFav) GetUserById(id int64) *favourites.User {
	MockF := MockFav{}
	for user, _ := range *MockF.GetFavouritesDB() {
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
	user := favourites.User{Id: 1, FirstName: "ALEx", LastName: "Man", Email: "somemail@example.com"}
	emptyDB[user] = []assets.Asset{}
	return &emptyDB
}
