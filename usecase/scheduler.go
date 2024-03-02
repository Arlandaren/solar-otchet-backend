package usecase

import (
	"server/models"
	"server/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func ScheduleTask(c *gin.Context) {
	var ScheduleTemplate models.ScheduleTemplate
	if err := c.ShouldBindJSON(&ScheduleTemplate); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var err error
	ScheduleTemplate.UserID, err = utils.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&ScheduleTemplate)
	c.JSON(200, gin.H{"status": "success", "message": "Task binded"})
}

func DeleteTask(c *gin.Context) {
	// Получаем идентификатор пользователя из запроса
	userID, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	// Получаем идентификатор шаблона из параметра запроса
	ScheduleID := c.Query("id")

	// Проверяем, принадлежит ли шаблон текущему пользователю
	var count int64
	models.DB.Model(&models.ScheduleTemplate{}).Where("id = ? AND user_id = ?", ScheduleID, userID).Count(&count)
	if count == 0 {
		// Если шаблон не принадлежит пользователю, возвращаем ошибку
		c.JSON(403, gin.H{"message": "you don't have permission to delete this schedule", "status": "error"})
		return
	}

	// Удаляем шаблон из базы данных
	if err := models.DB.Delete(&models.ScheduleTemplate{}, ScheduleID).Error; err != nil {
		// Обработка ошибки при удалении
		c.JSON(500, gin.H{"message": "error deleting schedule", "status": "error"})
		return
	}

	// Возвращаем успешный статус
	c.JSON(200, gin.H{"message": "schedule deleted", "status": "success"})
}

func CheckDate(c *gin.Context) {
	id, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	Schedules := GetAllSchedules(id)

	now := time.Now()
	// Цикл по расписаниям
	for i, schedule := range Schedules {
		if schedule.Date.Before(now) {
			id := Schedules[i].TemplateId
			template, err := utils.GetTemplateByID(id)
			if err != nil {
				c.JSON(500, gin.H{"status": "error", "message": "there aren't schedules where time has come"})
			}

			orgName := template.OrgName
			end := time.Now()
			interval := time.Duration(template.Interval) * 24 * time.Hour
			start := end.Add(-interval)

			errorr := utils.Getstats(c, orgName, start, end)
			if errorr != nil {
				c.JSON(500, gin.H{"err": err.Error()})
				return
			}

			c.JSON(200, gin.H{"message": "report generated", "id": template, "sched": schedule})
		}
	}
}
