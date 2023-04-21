package dao

import (
	"shouyindemo/models"
	"shouyindemo/utils"
)

type UserDao struct {
	BaseDao
}

var User = UserDao{}

//修改用户名
func (UserDao) UpdateUserName(name string, id int) error {
	err := utils.MySqlClient.DB.Model(&models.User{}).Where("id = ?", id).Update("name", name).Error
	return err
}

//分页查询用户列表
func (UserDao) GetUserPageList(index int, size int) ([]models.User, error) {
	results := make([]models.User, 0)
	err := utils.MySqlClient.DB.Model(&models.User{}).Where("age < ?", 20).Where("name LIKE ?", "%王").Limit(size).Offset(index * size).Find(&results).Error
	return results, err
}
