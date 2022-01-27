package assets

import (
	"GWI_assingment/platform2.0-go-challenge/logger"
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadAssets(filepath string) {
	var assets []Asset

	jsonFile, err := os.Open(filepath)
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
	logger.Log.Info("Successfully Opened assets.json")
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		logger.Log.Error(err.Error())
		os.Exit(1)
	}
	if err := json.Unmarshal(byteValue, &assets); err != nil {
		logger.Log.Error(err.Error())
		return
	}

	for asset_id, asset := range assets {
		tmpAsset := Asset{}

		switch asset.AssetType {
		case "chart":
			chart := Chart{}
			if err := json.Unmarshal(asset.Asset, &chart); err != nil {
				logger.Log.Error(err.Error())
				return
			}
			tmpAsset.Asset, _ = json.Marshal(chart)
			tmpAsset.AssetType = "chart"

		case "insight":
			insight := Insight{}
			if err := json.Unmarshal(asset.Asset, &insight); err != nil {
				logger.Log.Error(err.Error())
				return
			}
			tmpAsset.Asset, _ = json.Marshal(insight)
			tmpAsset.AssetType = "insight"

		case "audience":
			audience := Audience{}
			if err := json.Unmarshal(asset.Asset, &audience); err != nil {
				logger.Log.Error(err.Error())
				return
			}
			tmpAsset.AssetType = "audience"
			tmpAsset.Asset, _ = json.Marshal(audience)

		default:
			logger.Log.Error("unable to unmarshal JSON data or differentiate the type")

			// fmt.Println("")
		}
		tmpAsset.AssetID = AssetId(asset_id)
		tmpAsset.Description = asset.Description

		A := AsDB{}
		assetPtr := A.GetAssetDB()
		(*assetPtr)[AssetId(asset_id)] = tmpAsset

	}
	// A := AsDB{}
	// assetPtr := A.GetAssetDB()
	// for id, asset := range *assetPtr {
	// 	fmt.Printf("\nAssetID:%d --> %+v\n", id, asset)
	// }
}
