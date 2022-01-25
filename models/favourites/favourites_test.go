package favourites_test

import (
	"GWI_assingment/platform2.0-go-challenge/models/assets"
	"GWI_assingment/platform2.0-go-challenge/models/favourites"
	"GWI_assingment/platform2.0-go-challenge/models/users"
	"encoding/json"
	"fmt"
	"testing"
)

var MockFavourites = CreateFavouritesDB()

type MockFav struct {
	MFav *map[users.User][]assets.Asset
}

func (mf *MockFav) GetFavouritesDB() {
	mf.MFav = &MockFavourites
}

func TestGetFavouritesFromUserSuccess(t *testing.T) {
	F := MockFav{}
	F.GetFavouritesDB()
	resp := favourites.GetFavouritesFromUser(1)
	fmt.Println(resp)

}
func TestGetFavouritesFromUserFailure(t *testing.T) {

}
func TestAddFavouritesToUserSuccess(t *testing.T) {

}
func TestAddFavouritesToUserFailure(t *testing.T) {

}
func CreateFavouritesDB() map[users.User][]assets.Asset {
	result := map[users.User][]assets.Asset{}
	user := users.User{Id: 1, FirstName: "Test", LastName: "Man", Email: "somemail@example.com"}
	insight, _ := json.Marshal(assets.Insight{Text: "some text"})
	asset1 := assets.Asset{AssetType: "insight", Description: "hello insight", AssetID: 3, Asset: insight}
	audience, _ := json.Marshal(assets.Audience{Gender: "male", BirthCountry: "Greece", AgeGroup: 0, HoursInSocial: 3, Purchases: 20.0})
	asset2 := assets.Asset{AssetType: "audience", Description: "hello audience", AssetID: 4, Asset: audience}

	assetlist := []assets.Asset{}
	assetlist = append(assetlist, asset1, asset2)
	result[user] = assetlist

	return result
}
