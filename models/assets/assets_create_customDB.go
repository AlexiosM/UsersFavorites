package assets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

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

	for asset_id, asset := range assets {
		fmt.Println("-----------------------")
		tmpAsset := Asset{}

		switch asset.AssetType {
		case "chart":
			chart := Chart{}
			fmt.Printf("Id: %d\n", asset_id)
			if err := json.Unmarshal(asset.Asset, &chart); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%+v\n", chart)
			fmt.Printf("Description: %s\n", asset.Description)
			tmpAsset.AssetType = "chart"
			tmpAsset.Description = asset.Description
			AssetsDB[AssetId(asset_id)] = tmpAsset

		case "insight":
			insight := Insight{}
			fmt.Printf("Id: %d\n", asset_id)
			if err := json.Unmarshal(asset.Asset, &insight); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%+v\n", insight)
			fmt.Printf("Description: %s\n", asset.Description)
			tmpAsset.AssetType = "insight"
			tmpAsset.Description = asset.Description
			AssetsDB[AssetId(asset_id)] = tmpAsset

		case "audience":
			audience := Audience{}
			fmt.Printf("Id: %d\n", asset_id)
			if err := json.Unmarshal(asset.Asset, &audience); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%+v\n", audience)
			fmt.Printf("Description: %s\n", asset.Description)
			tmpAsset.AssetType = "audience"
			tmpAsset.Description = asset.Description
			AssetsDB[AssetId(asset_id)] = tmpAsset

		default:
			fmt.Println("unable to unmarshal JSON data or differentiate the type")
		}
	}
	fmt.Println(AssetsDB)
}
