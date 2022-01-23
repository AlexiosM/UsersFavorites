package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"GWI_assingment/platform2.0-go-challenge/models"
	"GWI_assingment/platform2.0-go-challenge/models/favourites"

	"github.com/gin-gonic/gin"
)

type Description struct {
	Description string `json:"description"`
}

func AddFavorites(c *gin.Context) {
	var favList favourites.ListOfFavourites

	id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id should be a number"})
		return
	}
	err = c.ShouldBind(&favList)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format of favourites"})
		return
	}

	resp := favourites.AddFavouritesToUser(id, favList)

	fmt.Println("Alex")
	fmt.Println(resp)
	if resp.Error != nil {
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

	resp := models.EditDescription(description.Description, id)
	if resp.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to edit DB"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id, "asset type": resp.AssetType, "new description": resp.Description})
}
