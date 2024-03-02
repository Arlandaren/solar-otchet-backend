package models

import "time"

type Histories struct {
	UserID      uint
	Orgname     string
	CurrentTime time.Time
	Start       time.Time
	End         time.Time
}
