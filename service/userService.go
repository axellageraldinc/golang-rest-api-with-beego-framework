package service

import (
	"axell_first_rest/models"
	"axell_first_rest/repository"
	"github.com/astaxie/beego/orm"
	"axell_first_rest/response"
)

func GetAllUser() map[string]*models.User {
	return repository.GetAllUser()
}

func GetOneUser(userId int) *response.UserResponse {
	var userResponse response.UserResponse
	user, e := repository.GetOneUser(userId)
	if e == orm.ErrNoRows {
		userResponse.Status = "No records found!"
	} else if e == orm.ErrMissPK {
		userResponse.Status = "No primary key found!"
	} else {
		userResponse.Status = "OK"
		userResponse.Name = user.Name
		userResponse.Age = user.Age
	}
	return &userResponse
}

func PostNewUser(user *models.User) *response.UserResponse {
	var userResponse response.UserResponse
	user, err := repository.PostNewUser(user)
	if err != nil {
		userResponse.Status = err.Error()
	} else {
		userResponse.Status = "OK"
		userResponse.Name = user.Name
		userResponse.Age = user.Age
	}
	return &userResponse
}

func UpdateExistingUser(userId int, newUserData *models.User) *response.UserResponse {
	var userResponse response.UserResponse
	_ , err := repository.UpdateExistingUser(userId, newUserData)
	if err != nil {
		userResponse.Status = err.Error()
	} else {
		userResponse.Status = "OK"
		userResponse.Name = newUserData.Name
		userResponse.Age = newUserData.Age
	}
	return &userResponse
}

func DeleteUser(userId int) *response.UserResponse {
	var userResponse response.UserResponse
	user, err := repository.DeleteUser(userId)
	if err == nil {
		userResponse.Status = "OK"
		userResponse.Name = user.Name
		userResponse.Age = user.Age
	} else {
		userResponse.Status = err.Error()
	}
	return &userResponse
}