package controller

import (
	"net/http"

	"github.com/DannyahIA/personal-server/internal/database"
	"github.com/DannyahIA/personal-server/internal/model"
	"github.com/gin-gonic/gin"
)

func ListRecentItemsHandler(c *gin.Context) {
	var recentItems []model.Recent
	err := database.Gorm.Table("recents").Find(&recentItems).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var itemIds []string
	for _, recent := range recentItems {
		itemIds = append(itemIds, recent.ItemId)
	}

	err = database.Gorm.Table("items").Where("item_id IN (?)", itemIds).Find(&Files).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if Files != nil {
		c.JSON(http.StatusOK, Files)
	} else {
		c.JSON(http.StatusOK, []model.Recent{})
	}
}

func AddRecentItemHandler(c *gin.Context) {
	var recent model.Recent
	err := c.BindJSON(&recent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = database.Gorm.Table("recents").Create(&recent).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, recent)
}

func DeleteRecentItemHandler(c *gin.Context) {
	itemId := c.Param("item_id")
	err := database.Gorm.Table("recents").Where("item_id = ?", itemId).Delete(&model.Recent{}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
