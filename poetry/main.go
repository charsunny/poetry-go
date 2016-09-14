package main

import (
	"fmt"
	_ "poetry/routers"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"

	"lib/ysqi/tokenauth2beego"

	"lib/ysqi/tokenauth2beego/o2o"

	"github.com/astaxie/beego/context"
)

func init() {

	beego.Debug("init main function")

	tokenauth2beego.Init("charsunny")

	orm.RegisterDriver("mysql", orm.DRMySQL)

	//orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:3306)/poetry?charset=utf8")
	username := beego.AppConfig.String("mysql::username")
	password := beego.AppConfig.String("mysql::password")
	host := beego.AppConfig.String("mysql::host")
	datebase := beego.AppConfig.String("mysql::datebase")
	connectStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, datebase)
	beego.Debug(connectStr)
	orm.RegisterDataBase("default", "mysql", connectStr)
}

func CheckFileter() beego.FilterFunc {
	return func(ctx *context.Context) {
		if token, err := o2o.Auth.CheckToken(ctx.Request); err == nil {
			ctx.Input.SetParam("uid", token.SingleID)
		}
	}
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.BConfig.WebConfig.StaticDir["/static"] = "static"
	beego.InsertFilter("/*", beego.BeforeRouter, CheckFileter())
	beego.InsertFilter("/v1/user/*", beego.BeforeRouter, o2o.DefaultFileter())
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
