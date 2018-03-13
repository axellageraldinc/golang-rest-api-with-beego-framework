package main

import (
	_ "axell_first_rest/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"axell_first_rest/models"
)

func init()  {
	dbUser := "root"
	dbName := "golang_dev"
	dbPwd := ""
	dbLink := fmt.Sprintf("%s:%s@/%s?charset=utf8", dbUser, dbPwd, dbName)

	orm.RegisterDriver(
		"mysql",
		orm.DRMySQL)
	orm.RegisterDataBase(
		"default",
		"mysql",
		dbLink)
	orm.RegisterModel(new(models.User))
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	// Database alias.
	name := "default"
	// Drop table and re-create.
	force := true
	// Print log.
	verbose := true
	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	beego.Run()
}
