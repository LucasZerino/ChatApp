package models

import (
	"gorm.io/gorm"
)

type WorkingHour struct {
	gorm.Model
	InboxID        int  `gorm:"not null"`
	AccountID      int  `gorm:"not null;index"`
	DayOfWeek      int  `gorm:"not null"`
	ClosedAllDay   bool `gorm:"default:false"`
	OpenHour       int  `gorm:"default:0"`
	OpenMinutes    int  `gorm:"default:0"`
	CloseHour      int  `gorm:"default:0"`
	CloseMinutes   int  `gorm:"default:0"`
	OpenAllDay     bool `gorm:"default:false"`
}

// Migrate the schema
func (WorkingHour) TableName() string {
	return "working_hours"
}
