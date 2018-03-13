package controllers

import (
	"github.com/astaxie/beego"
	"axell_first_rest/service"
	"axell_first_rest/models"
	"encoding/json"
)

type UserController struct {
	beego.Controller
}

func (userController *UserController) GetAll() {
	userList := service.GetAllUser()
	userController.Data["json"] = userList
	userController.ServeJSON()
}

func (userController *UserController) GetOneUser() {
	userId, _ := userController.GetInt(":userId", 0)
	user := service.GetOneUser(userId)
	userController.Data["json"] = user
	userController.ServeJSON()
}

func (userController *UserController) PostNewUser() {
	var newUser models.User
	json.Unmarshal(userController.Ctx.Input.RequestBody, &newUser)
	user := service.PostNewUser(&newUser)
	userController.Data["json"] = user
	userController.ServeJSON()
}

func (userController *UserController) UpdateExistingUser() {
	userId, _ := userController.GetInt(":userId", 0)
	var newUser models.User
	json.Unmarshal(userController.Ctx.Input.RequestBody, &newUser)
	user := service.UpdateExistingUser(userId, &newUser)
	userController.Data["json"] = user
	userController.ServeJSON()
}

func (userController *UserController) DeleteUser() {
	userId, _ := userController.GetInt(":userId", 0)
	user := service.DeleteUser(userId)
	userController.Data["json"] = user
	userController.ServeJSON()
}