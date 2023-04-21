package models

//数据库表结构
type User struct {
	Id    int    `gorm:"primaryKey;autoIncrement" form:"id"`
	Name  string `form:"name"`
	Phone int    `form:"phone"`
}

// 自定义表名
func (User) TableName() string {
	return "user"
}
