package models

// 表名
const (
	TableLocation = "tb_location"
)

// Location 用户地理位置.
type Location struct {
	ID        int64   `gorm:"primary_key;column:id"`
	UserID    int64   `gorm:"column:user_id"`
	Latitude  float64 `gorm:"column:latitude;type:DECIMAL(12, 10)"`
	Longitude float64 `gorm:"column:longitude;type:DECIMAL(13, 10)"`
	GeoHash   string  `gorm:"column:geo_hash;index"`

	TimeMixin
}

// TableName 表名.
func (Location) TableName() string {
	return TableLocation
}
