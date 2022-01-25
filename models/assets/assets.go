package assets

import (
	"encoding/json"
)

type AssetId int64

var AssetsDB = map[AssetId]Asset{}

type IAsDB interface {
	GetAssetDB()
}
type AsDB struct {
	As *map[AssetId]Asset
}

func (a *AsDB) GetAssetDB() {
	a.As = &AssetsDB
}

type Asset struct {
	AssetType   string `json:"asset_type"`
	AssetID     AssetId
	Description string `json:"description"`
	Asset       json.RawMessage
}

type Chart struct {
	SmallTitle string `json:"small_title"`
	Title      string `json:"title"`
	Data       string `json:"data"`
}

type Insight struct {
	Text string `json:"text"`
}

type Audience struct {
	Gender        string    `json:"gender"`
	BirthCountry  string    `json:"birth_country"`
	AgeGroup      AgeGroups `json:"age_group"`
	HoursInSocial int64     `json:"hours_in_social"`
	Purchases     float64   `json:"purchases"`
}

type AgeGroups int64

const (
	GenZ        AgeGroups = iota // 10-25
	Millennials                  // 26-41
	GenX                         // 42-57
	Boomers2                     // 58 - 67
	Boomers1                     // 68 – 76
	PostWar                      // 77 – 94
	WW2                          // 95 – 100
)

type AssetRespData struct {
	FavId     int64
	AssetType string
}

func (ag AgeGroups) GetAgeGroup() string {
	groups := []string{"GenZ", "Millennials", "GenX", "Boomers2", "Boomers1", "PostWar", "WW2"}
	return groups[ag]
}
