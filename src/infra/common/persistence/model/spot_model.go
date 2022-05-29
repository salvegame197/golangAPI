package model

import "gorm.io/gorm"

// Spot is the model of spots
type Spot struct {
	ID        string  `gorm:"primary_key;column:id;type:uuid" json:"id"`
	Name      string  `gorm:"size:100;not null;" json:"first_name"`
	Address   string  `gorm:"size:100;not null;" json:"address"`
	CreatedAt string  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt string  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

func (u *Spot) ToArray() map[string]interface{} {
	spot := map[string]interface{}{
		"ID":        u.ID,
		"Name":      u.Name,
		"Address":   u.Address,
		"CreatedAt": u.CreatedAt,
		"UpdatedAt": u.UpdatedAt,
	}
	return spot
}

func (u *Spot) Migrate(dbConnection *gorm.DB) {
	dbConnection.AutoMigrate(Spot{})
}
