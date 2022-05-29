package model

import "gorm.io/gorm"

type Tables struct {
	User User
	Spot Spot
}

func Initialize(dbConnection *gorm.DB) {
	tables := Tables{
		User: User{},
		Spot: Spot{},
	}
	tables.MigrateAll(dbConnection)
}

type Interface interface {
	MigrateAll(dbConnection *gorm.DB)
}

func (t *Tables) MigrateAll(dbConnection *gorm.DB) {
	t.Spot.Migrate(dbConnection)
	t.User.Migrate(dbConnection)

}
