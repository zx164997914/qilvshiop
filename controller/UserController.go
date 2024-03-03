package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	config "shouyindemo/conf"
	"shouyindemo/dao"
	"shouyindemo/models"
	"shouyindemo/service"
	"strconv"
)

type UserController struct {
}

func (controller *UserController) Add(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"msg": "参数错误"})
		return
	}
	dao.User.Add(&user)

	context.JSON(http.StatusOK, gin.H{"user": user})
}

//查询用户
func (controller *UserController) Get(context *gin.Context) {
	id, _ := strconv.Atoi(context.Query("id"))
	println("id:", id)

	user, err := service.GetUser(id)
	context.JSON(http.StatusOK, gin.H{
		"id":   id,
		"conf": config.GetConfig(),
		"user": user,
		"err":  err,
		//"hello": hello,
	})
}

//分页查询
func (controller *UserController) GetPageList(context *gin.Context) {
	index, _ := strconv.Atoi(context.Query("index"))
	println("index:", index)
	size, _ := strconv.Atoi(context.Query("size"))
	println("size:", size)

	list, _ := dao.User.GetUserPageList(index, size)
	context.JSON(http.StatusOK, gin.H{
		"users": list,
	})
}
func (controller *UserController) UpdateUserName(context *gin.Context) {
	name := context.PostForm("name")
	println("name:", name)
	id, _ := strconv.Atoi(context.PostForm("id"))
	println("id:", id)

	_ = dao.User.UpdateUserName(name, id)
	context.JSON(http.StatusOK, gin.H{
		"code": 1,
	})
}

func (controller *UserController) GetUserInfoById(context *gin.Context) {
	userId := context.GetInt("user_id")
	if userId == 0 {
		context.JSON(10000, "参数错误")
	}
	userInfo, err := service.GetUser(userId)
	if err != nil {

	}
	context.JSON(http.StatusOK, gin.H{
		"user": userInfo,
	})
}
