package monthmodel

import (
	"fetchAPI_gofiber/src/config"

	"gorm.io/gorm"
)

type Month struct {
	gorm.Model
	Name string
	Day  uint
}

func SelectAllMonth() []*Month {
	var items []*Month
	config.DB.Find(&items)
	return items
}

func SelectMonthById(id string) *Month {
	var item Month
	config.DB.First(&item, "id = ?", id)
	return &item
}

func PostMonth(item *Month) error {
	result := config.DB.Create(&item)
	return result.Error
}

func UpdateMonth(id int, newMonth *Month) error {
	var item Month
	result := config.DB.Model(&item).Where("id = ?", id).Updates(newMonth)
	return result.Error
}

func DeleteMonth(id int) error {
	var item Month
	result := config.DB.Delete(&item, "id = ?", id)
	return result.Error
}

func FindData(keyword string) []*Month {
	var items []*Month
	keyword = "%" + keyword + "%"
	config.DB.Where("CAST(id AS TEXT) LIKE ? OR name LIKE ? OR CAST(day AS TEXT) LIKE ?", keyword, keyword, keyword).Find(&items)
	return items
}

func FindCond(sort string, limit int, offset int) []*Month {
	var items []*Month
	config.DB.Order(sort).Limit(limit).Offset(offset).Find(&items)
	return items
}

func CountData() int64 {
	var count int64
	config.DB.Model(&Month{}).Count(&count)
	return count
}
