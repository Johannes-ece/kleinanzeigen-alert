package model

import (
	"github.com/jinzhu/gorm"
)

// Query that is beeing sored
type Query struct {
	gorm.Model
	ChatID   int64 `gorm:"index:chatid"`
	LastAds  []Ad
	Term     string `gorm:"type:varchar(100)"`
	Radius   int
	City     int
	MinPrice int
	MaxPrice int
	SaleType string `gorm:"type:varchar(100)"`
	CityName string `gorm:"type:varchar(100)"`
}

// AfterDelete delete all assiciated ads
func (u *Query) AfterDelete(tx *gorm.DB) (err error) {
	tx.Where("query_id = ?", u.ID).Delete(&Ad{})
	return
}
