package migrations

import (
	"time"

	"gorm.io/gorm"
)

type Alerts struct {
	ID          uint      `gorm:"primary key;autoIncrement" json: "id"`
	Type        *string   `json:"type"`
	Description *string   `json:"description"`
	Created     time.Time `json:"created"`
	Country     *string   `json:"country"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&Alerts{})
	return err
}
