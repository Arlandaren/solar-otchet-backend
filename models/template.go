package models

import "time"

type Template struct {
	UserID   uint
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null" json:"name"`
	OrgName  string `json:"orgname"`
	Interval uint16 `json:"interval"`
}

type ScheduleTemplate struct {
	Id     uint `gorm:"primaryKey"`
	UserID uint
	Date   time.Time `json:"date"` //формат "2023-12-11T13:02:16Z"
	// TemplateName string    `json:"templatename"`
	TemplateId uint `json:"templateid"`
}
