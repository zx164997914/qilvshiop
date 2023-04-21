package service

import (
	"shouyindemo/dao"
	"shouyindemo/models"
)

func GetUser(id int) (models.User, error) {
	var user models.User
	err := dao.User.Get(&user, id)
	if err != nil {
		return user, err
	}
	return user, nil
}
