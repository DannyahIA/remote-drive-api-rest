package controller

import (
	"net/http"

	"github.com/DannyahIA/personal-server/internal/database"
	"github.com/DannyahIA/personal-server/internal/model"
	"github.com/gin-gonic/gin"
)

var Files []model.Item

func ListItemsHandler(c *gin.Context) {
	var items []model.Item
	err := database.Gorm.Table("items").Find(&items).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if items != nil {
		c.JSON(http.StatusOK, items)
	} else {
		c.JSON(http.StatusOK, []model.Item{})
	}
}

func CreateItemHandler(c *gin.Context) {
	var item model.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.Gorm.Table("items").Create(&item).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func DeleteItemHandler(c *gin.Context) {
	var requestBody []struct {
		ID uint `json:"id"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, delete := range requestBody {
		id := delete.ID

		if id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id query parameter is required"})
			return
		}

		err := database.Gorm.Table("items").Where("id = ?", id).Delete(&model.Item{}).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.Status(http.StatusOK)
}

func UpdateItemHandler(c *gin.Context) {
	var requestBody []struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, req := range requestBody {
		id := req.ID
		name := req.Name

		if id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id query parameter is required"})
			return
		}

		if name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "name query parameter is required"})
			return
		}

		err := database.Gorm.Table("items").Where("id = ?", id).Update("name", name).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Items updated successfully"})
}

func SearchItemHandler(c *gin.Context) {
	var requestBody struct {
		Query string `json:"query"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := requestBody.Query
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query query parameter is required"})
		return
	}

	var items []model.Item
	err := database.Gorm.Table("items").Where("name LIKE ?", "%"+query+"%").Find(&items).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if items != nil {
		c.JSON(http.StatusOK, items)
	} else {
		c.JSON(http.StatusOK, []model.Item{})
	}
}

func UploadItemHandler(c *gin.Context) {
	var requestBody struct {
		Items []model.Item `json:"items"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, item := range requestBody.Items {
		err := database.Gorm.Table("items").Create(&item).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.Status(http.StatusCreated)
}

func DownloadItemHandler(c *gin.Context) {
	var requestBody struct {
		ID uint `json:"id"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := requestBody.ID
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id query parameter is required"})
		return
	}

	var item model.Item
	err := database.Gorm.Table("items").Where("id = ?", id).Find(&item).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}