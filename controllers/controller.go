package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ChangeDescription(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("asset_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": strconv.FormatInt(int64(id), 10)})
	}
	//models.EditDescription(id)
	c.JSON(http.StatusOK, gin.H{"message": "edited"})
}
