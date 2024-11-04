package controller

import (
	"github.com/DannyahIA/personal-server/internal/database"
	"github.com/DannyahIA/personal-server/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LoginHandler(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")

	if email == "" && password == "" {
		c.JSON(400, gin.H{
			"message": "email and password are required",
		})
		return
	}

	var user model.Profile

	err := database.Gorm.
		Table("users").
		Where("email = ? AND password_hash = ?", email, password).
		First(&user).
		Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": "email or password is incorrect",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "login success",
		"data":    user,
	})

}

func RegisterHandler(c *gin.Context) {
	var user model.Profile

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request",
		})
		return
	}

	uuid, _ := uuid.NewV7()
	user.UserId = uuid.String()

	err = database.Gorm.
		Table("users").
		Create(&user).
		Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": "failed to register",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func GetProfileHandler(c *gin.Context) {
	var user model.Profile

	err := database.Gorm.
		Table("users").
		First(&user).
		Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": "user not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    user,
	})
}

func UpdateProfileHandler(c *gin.Context) {
	var user model.Profile

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request",
		})
		return
	}

	err = database.Gorm.
		Table("users").
		Where("user_id = ?", user.UserId).
		Updates(user).
		Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": "failed to update profile",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func DeleteProfileHandler(c *gin.Context) {
	user_id := c.Query("user_id")

	err := database.Gorm.
		Table("users").
		Where("user_id = ?", user_id).
		Delete(model.Profile{}).
		Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": "failed to delete profile",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func ForgotPasswordHandler(c *gin.Context) {
	email := c.Query("email")

	if email == "" {
		c.JSON(400, gin.H{
			"message": "email is required",
		})
		return
	}

	var user model.Profile

	err := database.Gorm.
		Table("users").
		Where("email = ?", email).
		First(&user).
		Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": "user not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success", //TODO: send email with link to new password
	})
}
