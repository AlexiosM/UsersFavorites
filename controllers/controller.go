package controllers

import (
	"net/http"
	"strconv"

	"GWI_assingment/platform2.0-go-challenge/models/assets"
	"GWI_assingment/platform2.0-go-challenge/models/favourites"
	"GWI_assingment/platform2.0-go-challenge/models/utils"

	"github.com/gin-gonic/gin"
)

type Description struct {
	Description string `json:"description"`
}

func GetFavourites(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id should be a number"})
		return
	}
	var favouritesDB favourites.Fav // This is the structure to be mocked for testing
	resp := favourites.GetFavouritesFromUser(id, &favouritesDB)
	if resp.Error != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": resp.Error})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func AddFavourites(c *gin.Context) {
	var incomingFavList favourites.ListOfFavourites

	id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id should be a number"})
		return
	}
	err = c.ShouldBind(&incomingFavList)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format of favourites"})
		return
	}

	var favouritesDB favourites.Fav // This is the structure to be mocked for testing
	var assetsDB assets.AsDB
	resp := favourites.AddFavouritesToUser(id, incomingFavList, &favouritesDB, &assetsDB)

	if resp.Error != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add to favourites"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ChangeDescription(c *gin.Context) {
	var description Description

	id, err := strconv.ParseInt(c.Param("asset_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id should be a number"})
		return
	}
	err = c.ShouldBind(&description)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "description should be given"})
		return
	}

	var favouritesDB favourites.Fav // This is the structure to be mocked for testing
	var assetsDB assets.AsDB
	resp := utils.EditAssetDescription(description.Description, assets.AssetId(id), &favouritesDB, &assetsDB)
	if resp.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to edit DB"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id, "asset type": resp.AssetType, "new description": resp.Description})
}
