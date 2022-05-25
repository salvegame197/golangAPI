package config

func RunMigration() {
	DBmap.AutoMigrate(&User{})
}
