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
			if err := json.Unmarshal(asset.Asset, &chart); err != nil {
				fmt.Println(err)
				return
			}
			tmpAsset.Asset, _ = json.Marshal(chart)
			tmpAsset.AssetType = "chart"
			tmpAsset.AssetID = AssetId(asset_id)
			tmpAsset.Description = asset.Description
			AssetsDB[AssetId(asset_id)] = tmpAsset

		case "insight":
			insight := Insight{}
			if err := json.Unmarshal(asset.Asset, &insight); err != nil {
				fmt.Println(err)
				return
			}
			tmpAsset.Asset, _ = json.Marshal(insight)
			tmpAsset.AssetType = "insight"
			tmpAsset.AssetID = AssetId(asset_id)
			tmpAsset.Description = asset.Description
			AssetsDB[AssetId(asset_id)] = tmpAsset

		case "audience":
			audience := Audience{}
			if err := json.Unmarshal(asset.Asset, &audience); err != nil {
				fmt.Println(err)
				return
			}
			tmpAsset.AssetType = "audience"
			tmpAsset.Asset, _ = json.Marshal(audience)
			tmpAsset.AssetID = AssetId(asset_id)
			tmpAsset.Description = asset.Description
			AssetsDB[AssetId(asset_id)] = tmpAsset

		default:
			fmt.Println("unable to unmarshal JSON data or differentiate the type")
		}

	}
	for id, asset := range AssetsDB {
		fmt.Printf("\nAssetID:%d --> %+v\n", id, asset)
	}
}
