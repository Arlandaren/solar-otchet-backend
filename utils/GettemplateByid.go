package utils

import (
	"server/models"
)

func GetTemplateByID(id uint) (*models.Template, error) {
	var template models.Template

	//поиска записи в таблице Template по ID
	err := models.DB.First(&template, id).Error
	if err != nil {
		// Обработка ошибки (например, запись не найдена)
		return nil, err
	}

	// Вернуть найденную запись
	return &template, nil
}
