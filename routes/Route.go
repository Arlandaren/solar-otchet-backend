package routes

import (
	"server/controllers"
	"server/middlwares"
	"server/usecase"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {

	api := r.Group("/api/v1")
	{
		api.Use(middlwares.AuthMiddleware())

		users := api.Group("/user")
		{
			users.GET("/getuserinfo", usecase.SendAllUserInfo)
		}
		stats := api.Group("stats/")
		{
			stats.GET("/report", usecase.Getstats)
			stats.GET("/download", utils.DownloadMinioFile)
		}

		scheduler := api.Group("/scheduler")
		{
			scheduler.POST("/createtask", usecase.ScheduleTask)
			scheduler.GET("/check", usecase.CheckDate)
			scheduler.DELETE("/deletetask", usecase.DeleteTask)
		}

		templates := api.Group("/template")
		{
			templates.POST("/save", usecase.Gettemplate)
			templates.DELETE("/delete", usecase.DeleteTemplate)
		}
	}
	auth := r.Group("/auth")
	{
		auth.POST("signup", controllers.Signup)
		auth.POST("login", controllers.Login)
	}
}
