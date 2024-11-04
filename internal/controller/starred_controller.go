package controller

import (
	"net/http"

	"github.com/DannyahIA/personal-server/internal/database"
	"github.com/DannyahIA/personal-server/internal/model"
	"github.com/gin-gonic/gin"
)

func ListStarredItemsHandler(c *gin.Context) {
	var starredItems []model.Starred
	err := database.Gorm.Table("starred").Find(&starredItems).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var itemIds []string
	for _, starred := range starredItems {
		itemIds = append(itemIds, starred.ItemId)
	}

	err = database.Gorm.Table("items").Where("item_id IN (?)", itemIds).Find(&Files).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if Files != nil {
		c.JSON(http.StatusOK, Files)
	} else {
		c.JSON(http.StatusOK, []model.Starred{})
	}
}

func StarItemHandler(c *gin.Context) {
	var starred model.Starred
	err := c.BindJSON(&starred)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = database.Gorm.Table("starred").Create(&starred).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, starred)
}

func UnstarItemHandler(c *gin.Context) {
	itemId := c.Param("item_id")
	err := database.Gorm.Table("starred").Where("item_id = ?", itemId).Delete(model.Starred{}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}