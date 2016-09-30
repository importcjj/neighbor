package models

// 表名
const (
	TableUser = "tb_user"
)

// User 用户信息.
type User struct {
	ID       int64  `gorm:"primary_key;column:id"`
	Username string `gorm:"column:username"`
	Sex      string `gorm:"column:sex"`

	TimeMixin
}

// TableName 表名.
func (User) TableName() string {
	return TableUser
}
