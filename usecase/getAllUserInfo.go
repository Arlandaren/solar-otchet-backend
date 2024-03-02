package usecase

import (
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
)

// GetAllSchedules получает все расписания пользователя по его идентификатору
func GetAllSchedules(id uint) []models.ScheduleTemplate {
	var ScheduleTemplate []models.ScheduleTemplate
	err := models.DB.Where("user_id = ?", id).Find(&ScheduleTemplate).Error
	if err != nil {
		return nil
	}
	return ScheduleTemplate
}

// GetAllTemplates получает все шаблоны пользователя по его идентификатору
func GetAllTemplates(id uint) []models.Template {
	var Template []models.Template
	err := models.DB.Where("user_id = ?", id).Find(&Template).Error
	if err != nil {
		return nil
	}
	return Template
}
func GetAllGeneratedStats(id uint) []models.GeneratedStats {
	var GeneratedStats []models.GeneratedStats
	err := models.DB.Where("user_id = ?", id).Find(&GeneratedStats).Error
	if err != nil {
		return nil
	}
	return GeneratedStats
}

func SendAllUserInfo(c *gin.Context) {
	// Получаем идентификатор пользователя
	id, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	// Инициализируем объект пользователя
	var User models.User

	// Поиск пользователя в базе данных по его идентификатору
	err = models.DB.Where("id = ?", id).Find(&User).Error
	if err != nil {
		// Обработка ошибки при поиске пользователя
		c.JSON(500, gin.H{"message": "error with connecting to database", "status": "error"})
		return
	}

	// Получаем все расписания пользователя
	schedules := GetAllSchedules(User.ID)

	// Получаем все шаблоны пользователя
	templates := GetAllTemplates(User.ID)

	//получаем все сгенерированные отчеты
	generatedStats := GetAllGeneratedStats(User.ID)

	// Присваиваем расписания и шаблоны объекту пользователя
	User.Schedules = schedules
	User.Templates = templates
	User.GeneratedStats = generatedStats

	// Отправляем полную информацию о пользователе в формате JSON
	c.JSON(200, User)
}
