package controller

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/DannyahIA/personal-server/internal/database"
	"github.com/DannyahIA/personal-server/internal/model"
	"github.com/DannyahIA/personal-server/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ListBackupItemsHandler(c *gin.Context) {
	var backupItems []model.Backup
	err := database.Gorm.Table("backup").Find(&backupItems).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if backupItems != nil {
		c.JSON(http.StatusOK, backupItems)
	} else {
		c.JSON(http.StatusOK, []model.Backup{})
	}
}

func GetBackupItemHandler(c *gin.Context) {
	backupId := c.Query("backup_id")

	var backup model.Backup
	err := database.Gorm.Table("backup").Where("backup_id = ?", backupId).First(&backup).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, backup)
}

func CreateNewBackupHandler(c *gin.Context) {
	userId := c.Query("user_id")

	bkpName := "Backup-" + time.Now().Format("02-01-2006")
	zipFilePath := filepath.Join(".", "backup", userId, bkpName+".zip")

	err := util.ZipFolder(filepath.Join(".", "drive", userId), zipFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		os.Remove(zipFilePath)
		return
	}

	uuid, _ := uuid.NewV7()
	backup := model.Backup{
		BackupId:  uuid.String(),
		UserId:    userId,
		Name:      bkpName,
		Path:      zipFilePath,
		CreatedAt: time.Now(),
	}

	err = database.Gorm.Debug().Table("backup").Create(&backup).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		os.Remove(zipFilePath)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Backup created successfully"})
}

func RestoreBackupHandler(c *gin.Context) {
	backupId := c.Query("backup_id")

	var backup model.Backup

	err := database.Gorm.Table("backup").Where("backup_id = ?", backupId).First(&backup).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = util.UnzipFolder(backup.Path, filepath.Join(".", "drive"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Backup restored successfully"})
}

func DeleteBackupHandler(c *gin.Context) {
	backupId := c.Query("backup_id")

	var backup model.Backup

	err := database.Gorm.Table("backup").Where("backup_id = ?", backupId).First(&backup).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = os.Remove(backup.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = database.Gorm.Table("backup").Where("backup_id = ?", backupId).Delete(&backup).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Backup deleted successfully"})
}
