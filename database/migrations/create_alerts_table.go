package migrations

import (
	"time"

	"gorm.io/gorm"
)

type Alerts struct {
	ID          uint      `gorm:"primary key;autoIncrement" json: "id"`
	Type        *string   `json:"name"`
	Description *string   `json:"email"`
	Created     time.Time `json:"date"`
	Country     *string   `json:"country"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&Alerts{})
	return err
}
