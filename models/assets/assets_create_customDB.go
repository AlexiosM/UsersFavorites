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
	fmt.Println("Successfully Opened assets.json")
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Cound not read json file")
		os.Exit(1)
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

		case "insight":
			insight := Insight{}
			if err := json.Unmarshal(asset.Asset, &insight); err != nil {
				fmt.Println(err)
				return
			}
			tmpAsset.Asset, _ = json.Marshal(insight)
			tmpAsset.AssetType = "insight"

		case "audience":
			audience := Audience{}
			if err := json.Unmarshal(asset.Asset, &audience); err != nil {
				fmt.Println(err)
				return
			}
			tmpAsset.AssetType = "audience"
			tmpAsset.Asset, _ = json.Marshal(audience)

		default:
			fmt.Println("unable to unmarshal JSON data or differentiate the type")
		}
		tmpAsset.AssetID = AssetId(asset_id)
		tmpAsset.Description = asset.Description

		A := AsDB{}
		assetPtr := A.GetAssetDB()
		(*assetPtr)[AssetId(asset_id)] = tmpAsset

	}
	A := AsDB{}
	assetPtr := A.GetAssetDB()
	for id, asset := range *assetPtr {
		fmt.Printf("\nAssetID:%d --> %+v\n", id, asset)
	}
}
