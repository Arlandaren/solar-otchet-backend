package usecase

import (
	"fmt"
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func Gettemplate(c *gin.Context) {
	// Инициализация переменной Template для хранения данных из JSON-запроса
	var Template models.Template

	// Попытка привязать данные JSON-запроса к структуре Template
	if err := c.ShouldBindJSON(&Template); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Получение идентификатора пользователя с помощью функции GetUserId из утилиты utils
	var err error
	Template.UserID, err = utils.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	// Проверка существования записи с указанным именем в таблице базы данных
	var count int64
	models.DB.Model(&models.Template{}).Where("name = ?", Template.Name).Count(&count)

	// Если запись существует, создать уникальное имя
	if count > 0 {
		for i := 1; ; i++ {
			newName := fmt.Sprintf("%s (%d)", Template.Name, i)

			// Проверка существования записи с новым именем
			var excount int64
			models.DB.Model(&models.Template{}).Where("name = ?", newName).Count(&excount)

			// Если запись с новым именем не существует, присвоить его Template.Name и создать запись
			if excount == 0 {
				Template.Name = newName
				models.DB.Create(&Template)
				break
			}
		}
	} else {
		// Если запись с указанным именем не существует, просто создать запись
		models.DB.Create(&Template)
	}

	// Вернуть успешный JSON-ответ
	c.JSON(200, gin.H{"message": "template saved", "status": "success"})
}

func DeleteTemplate(c *gin.Context) {
	// Получаем идентификатор пользователя из запроса
	userID, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	// Получаем идентификатор шаблона из параметра запроса
	templateID := c.Query("id")

	// Проверяем, принадлежит ли шаблон текущему пользователю
	var count int64
	models.DB.Model(&models.Template{}).Where("id = ? AND user_id = ?", templateID, userID).Count(&count)
	if count == 0 {
		// Если шаблон не принадлежит пользователю, возвращаем ошибку
		c.JSON(403, gin.H{"message": "you don't have permission to delete this template", "status": "error"})
		return
	}

	// Удаляем шаблон из базы данных
	if err := models.DB.Delete(&models.Template{}, templateID).Error; err != nil {
		// Обработка ошибки при удалении
		c.JSON(500, gin.H{"message": "error deleting template", "status": "error"})
		return
	}

	// Возвращаем успешный статус
	c.JSON(200, gin.H{"message": "template deleted", "status": "success"})
}
