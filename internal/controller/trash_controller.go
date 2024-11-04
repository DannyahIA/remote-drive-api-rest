package controller

import (
	"net/http"

	"github.com/DannyahIA/personal-server/internal/database"
	"github.com/DannyahIA/personal-server/internal/model"
	"github.com/gin-gonic/gin"
)

func ListTrashItemsHandler(c *gin.Context) {
	var trashItems []model.Trash
	err := database.Gorm.Table("trash").Find(&trashItems).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var itemIds []string
	for _, trash := range trashItems {
		itemIds = append(itemIds, trash.ItemId)
	}

	err = database.Gorm.Table("items").Where("item_id IN (?)", itemIds).Find(&Files).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if Files != nil {
		c.JSON(http.StatusOK, Files)
	} else {
		c.JSON(http.StatusOK, []model.Trash{})
	}
}

func MoveToTrashHandler(c *gin.Context) {
	var item model.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	trash := model.Trash{
		ItemId: item.ItemId,
	}

	err := database.Gorm.Table("items").Where("item_id = ?", item.ItemId).Delete(&item).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = database.Gorm.Table("trash").Create(&trash).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item moved to trash successfully"})
}

func RestoreTrashItemHandler(c *gin.Context) {
	var trash model.Trash
	if err := c.ShouldBindJSON(&trash); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.Gorm.Table("trash").Where("item_id = ?", trash.ItemId).Delete(&trash).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item restored successfully"})
}

func DeleteTrashItemHandler(c *gin.Context) {
	var trash model.Trash
	if err := c.ShouldBindJSON(&trash); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.Gorm.Table("trash").Where("item_id = ?", trash.ItemId).Delete(&trash).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}

func EmptyTrashHandler(c *gin.Context) {
	err := database.Gorm.Table("trash").Delete(&model.Trash{}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Trash emptied successfully"})
}

