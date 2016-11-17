package main

import (
	_ "Coop-Agile/back-end/routers"
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	_"Coop-Agile/back-end/models"
	_ "github.com/lib/pq"
)

func main() {
	OrmName := "default"
	DatabaseType := "postgres"
	userName := "coop"
	password := "qweqwe123"
	dbname := "hacker"
	sslMode := "disable"

	beego.Info("____________________________________")
	beego.Info("Database Process Started.")
	orm.Debug = true
	sqlConnectionStatement := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=%v", userName, password, dbname, sslMode)
	orm.RegisterDataBase(OrmName, DatabaseType, sqlConnectionStatement)
	err := orm.RunSyncdb(OrmName, false, orm.Debug)
	if err != nil {
		beego.Error(err)

	}
	beego.Run()
}

