package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"unique" json:"name"`
	Password string `json:"password"`
	// Role      string             `json:"role"`	решено убрать до лучших времен
	Templates      []Template         `gorm:"foreignkey:UserID"`
	Schedules      []ScheduleTemplate `gorm:"foreignkey:UserID"`
	GeneratedStats []GeneratedStats   `gorm:"foreignkey:UserID"`
	History        []Histories        `gorm:"foreignkey:UserID"`
}
