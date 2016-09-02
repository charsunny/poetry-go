package main

import (
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
	orm.RegisterDataBase("default", "mysql", "r72g48k7ib:Sun@1989@tcp(rdsc580646shskz416s0.mysql.rds.aliyuncs.com:3306)/r72g48k7ib?charset=utf8")
}

func CheckFileter() beego.FilterFunc {
	return func(ctx *context.Context) {
		if token, err := o2o.Auth.CheckToken(ctx.Request); err == nil {
			ctx.Input.SetParam("uid",token.SingleID)
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
