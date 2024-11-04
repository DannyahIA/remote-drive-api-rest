package router

import (
	"github.com/DannyahIA/personal-server/internal/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AddAllowHeaders("*")
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "DELETE"}
	r.Use(cors.New(config))

	r.GET("/ping", controller.Ping)

	file_manager := r.Group("/file-manager")
	{
		itemRoute := file_manager.Group("/item")
		{
			itemRoute.GET("/list", controller.ListItemsHandler)
			itemRoute.POST("/new-item", controller.CreateItemHandler)
			itemRoute.DELETE("/", controller.DeleteItemHandler)
			itemRoute.PATCH("/", controller.UpdateItemHandler)
			itemRoute.GET("/search", controller.SearchItemHandler)
			itemRoute.GET("/download", controller.DownloadItemHandler)
			itemRoute.POST("/upload", controller.UploadItemHandler)
		}

		profileRoute := file_manager.Group("/profile")
		{
			profileRoute.GET("/login", controller.LoginHandler)
			profileRoute.POST("/register", controller.RegisterHandler)
			profileRoute.PATCH("/update", controller.UpdateProfileHandler)
			profileRoute.DELETE("/delete", controller.DeleteProfileHandler)
			profileRoute.GET("/forgot-password", controller.ForgotPasswordHandler)
		}

		backupRoute := file_manager.Group("/backup")
		{
			backupRoute.GET("/list", controller.ListBackupItemsHandler)
			backupRoute.GET("/", controller.GetBackupItemHandler)
			backupRoute.POST("/new-backup", controller.CreateNewBackupHandler)
			backupRoute.POST("/restore-backup", controller.RestoreBackupHandler)
			backupRoute.DELETE("/", controller.DeleteBackupHandler)
		}

		starredRoute := file_manager.Group("/starred")
		{
			starredRoute.GET("/list", controller.ListStarredItemsHandler)
			starredRoute.POST("/", controller.StarItemHandler)
			starredRoute.DELETE("/", controller.UnstarItemHandler)
		}

		recentRoute := file_manager.Group("/recent")
		{
			recentRoute.GET("/list", controller.ListRecentItemsHandler)
			recentRoute.POST("/", controller.AddRecentItemHandler)
			recentRoute.DELETE("/", controller.DeleteRecentItemHandler)
		}

		trashRoute := file_manager.Group("/trash")
		{
			trashRoute.GET("/list", controller.ListTrashItemsHandler)
			trashRoute.POST("/", controller.MoveToTrashHandler)
			trashRoute.POST("/restore", controller.RestoreTrashItemHandler)
			trashRoute.DELETE("/", controller.DeleteTrashItemHandler)
		}
	}
	return r
}
