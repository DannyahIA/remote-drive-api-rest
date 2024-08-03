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
		file_manager.GET("/list-root-folders", controller.ListRootFoldersHandler)
		file_manager.POST("/list-files", controller.ListFilesHandler)
		file_manager.GET("search-files", controller.SearchFileHandler)
		file_manager.POST("/create-folder/:folder", controller.CreateFolderHandler)
		file_manager.POST("/upload-file", controller.UploadFileHandler)
		file_manager.DELETE("/delete-file", controller.DeleteFileHandler)
		file_manager.GET("/download-file", controller.DownloadFileHandler)
		file_manager.PUT("/rename-file", controller.RenameFileHandler)
	}
	return r
}
