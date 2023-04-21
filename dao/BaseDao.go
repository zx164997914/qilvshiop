package dao

import (
	"errors"
	"fmt"
	"shouyindemo/utils"
)

type BaseDao struct {
}

//新增
func (BaseDao) Add(model interface{}) (err error) {
	err = utils.MySqlClient.DB.Create(model).Error
	return
}

//根据Id查询
func (BaseDao) Get(model interface{}, id int) error {
	fmt.Println("UserInfoDao Get")
	if id < 1 {
		return errors.New("请输入id")
	}
	_ = utils.MySqlClient.DB.First(&model, id)
	return nil
}

//更新
func (BaseDao) UpdateModel(value interface{}) (err error) {
	err = utils.MySqlClient.DB.Save(value).Error
	return
}

//删除
func (BaseDao) DeleteModel(value interface{}) (err error) {
	err = utils.MySqlClient.DB.Delete(value).Error
	return
}
