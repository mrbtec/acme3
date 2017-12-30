package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/auth"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/mrbtec/acme3/models"
	_ "github.com/mrbtec/acme3/routers"
)

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "acme.db")
	name := "default"
	force := false
	verbose := false
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.InsertFilter("/appadmin/*", beego.BeforeRouter, auth.Basic("youradminname", "YourAdminPassword"))
	beego.Run()
}
