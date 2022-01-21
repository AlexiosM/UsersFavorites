package assets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	assetsDB map[int64]Asset
)

type IAsset interface {
	EditAssetDescription(newDescription string) error
}

type Asset struct {
	AssetType   string `json:"asset_type"`
	Description string `json:"description"`
	Asset       json.RawMessage
}

type Chart struct {
	SmallTitle string `json:"small_title"`
	Title      string `json:"title"`
	Data       string `json:"data"`
	//Description string `json:"description"`
}

// func (c *Chart) EditAssetDescription(newDescription string) {
// 	c.Description = newDescription
// }

type Insight struct {
	Text string `json:"text"`
	//Description string `json:"description"`
}

// func (c *Insight) EditAssetDescription(newDescription string) {
// 	c.Description = newDescription
// }

type Audience struct {
	Gender        string    `json:"gender"`
	BirthCountry  string    `json:"birth_country"`
	AgeGroup      AgeGroups `json:"age_group"`
	HoursInSocial int64     `json:"hours_in_social"`
	Purchases     float64   `json:"purchases"`
	//Description   string    `json:"description"`
}

// func (c *Audience) EditAssetDescription(newDescription string) {
// 	c.Description = newDescription
// }

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

func (ag AgeGroups) GetAgeGroup() string {
	groups := []string{"GenZ", "Millennials", "GenX", "Boomers2", "Boomers1", "PostWar", "WW2"}
	return groups[ag]
}

func LoadAssets(filepath string) {
	var assets []Asset

	jsonFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Cound not read json file")
	}
	if err := json.Unmarshal(byteValue, &assets); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("================================")

	for _, asset := range assets {
		fmt.Println("-----------------------")
		switch asset.AssetType {
		case "chart":
			chart := Chart{}
			if err := json.Unmarshal(asset.Asset, &chart); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%+v\n", chart)

		case "insight":
			insight := Insight{}
			if err := json.Unmarshal(asset.Asset, &insight); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%+v\n", insight)

		case "audience":
			audience := Audience{}
			if err := json.Unmarshal(asset.Asset, &audience); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%+v\n", audience)
		default:
			fmt.Println("unable to unmarshal JSON data or differentiate the type")
		}
	}
}
