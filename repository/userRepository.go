package repository

import (
	"axell_first_rest/models"
	_ "github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/orm"
)

func GetAllUser() (list []orm.Params) {
	ormObject := orm.NewOrm()
	num, err := ormObject.Raw("SELECT name, age FROM user").Values(&list)
	if err == nil && num > 0 {
		return list
	} else {
		return nil
	}
}

func GetOneUser(userId int) (user models.User, err error) {
	ormObject := orm.NewOrm()
	user = models.User{Id:userId}
	err = ormObject.Read(&user)
	return user, err
}

func PostNewUser(user *models.User) (insertedUser *models.User, err error) {
	ormObject := orm.NewOrm()
	ormObject.Begin()
	_ , err = ormObject.Insert(user)
	if err == nil {
		ormObject.Commit()
		return user, nil
	} else {
		ormObject.Rollback()
		return nil, err
	}
}

func UpdateExistingUser(userId int, newUserData *models.User) (user models.User, err error) {
	ormObject := orm.NewOrm()
	ormObject.Begin()
	user = models.User{Id:userId}
	err = ormObject.Read(&user)
	if err == nil {
		user.Name = newUserData.Name
		user.Age = newUserData.Age
		_ , err = ormObject.Update(&user)
		if err == nil {
			ormObject.Commit()
		} else {
			ormObject.Rollback()
		}
	}
	return user, err
}

func DeleteUser(userId int) (user models.User, err error) {
	ormObject := orm.NewOrm()
	ormObject.Begin()
	user = models.User{Id:userId}
	err = ormObject.Read(&user)
	if err == nil {
		_, err = ormObject.Delete(&user)
		if err == nil {
			ormObject.Commit()
		} else {
			ormObject.Rollback()
		}
	}
	return user, err
}