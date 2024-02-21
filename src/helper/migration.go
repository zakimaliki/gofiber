package helper

import (
	"fetchAPI_gofiber/src/config"
	monthmodel "fetchAPI_gofiber/src/models/MonthModel"
)

func Migration() {
	config.DB.AutoMigrate(&monthmodel.Month{})
}
