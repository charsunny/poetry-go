package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["poetry/controllers:BaseController"] = append(beego.GlobalControllerRouter["poetry/controllers:BaseController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["poetry/controllers:FeedController"] = append(beego.GlobalControllerRouter["poetry/controllers:FeedController"],
		beego.ControllerComments{
			Method: "GetFeeds",
			Router: `/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["poetry/controllers:ObjectController"] = append(beego.GlobalControllerRouter["poetry/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["poetry/controllers:ObjectController"] = append(beego.GlobalControllerRouter["poetry/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["poetry/controllers:ObjectController"] = append(beego.GlobalControllerRouter["poetry/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["poetry/controllers:ObjectController"] = append(beego.GlobalControllerRouter["poetry/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["poetry/controllers:ObjectController"] = append(beego.GlobalControllerRouter["poetry/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["poetry/controllers:PoemController"] = append(beego.GlobalControllerRouter["poetry/controllers:PoemController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:pid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["poetry/controllers:PoemController"] = append(beego.GlobalControllerRouter["poetry/controllers:PoemController"],
		beego.ControllerComments{
			Method: "GetComments",
			Router: `/comments`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["poetry/controllers:PoemController"] = append(beego.GlobalControllerRouter["poetry/controllers:PoemController"],
		beego.ControllerComments{
			Method: "LikePoem",
			Router: `/like`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["poetry/controllers:RecommandController"] = append(beego.GlobalControllerRouter["poetry/controllers:RecommandController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/info`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["poetry/controllers:RecommandController"] = append(beego.GlobalControllerRouter["poetry/controllers:RecommandController"],
		beego.ControllerComments{
			Method: "GetList",
			Router: `/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["poetry/controllers:UserController"] = append(beego.GlobalControllerRouter["poetry/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["poetry/controllers:UserController"] = append(beego.GlobalControllerRouter["poetry/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["poetry/controllers:UserController"] = append(beego.GlobalControllerRouter["poetry/controllers:UserController"],
		beego.ControllerComments{
			Method: "UploadPic",
			Router: `/pic`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["poetry/controllers:UserController"] = append(beego.GlobalControllerRouter["poetry/controllers:UserController"],
		beego.ControllerComments{
			Method: "AddFeed",
			Router: `/addfeed`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["poetry/controllers:UserController"] = append(beego.GlobalControllerRouter["poetry/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetFeeds",
			Router: `/feeds`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["poetry/controllers:UserController"] = append(beego.GlobalControllerRouter["poetry/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["poetry/controllers:UserController"] = append(beego.GlobalControllerRouter["poetry/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetInfo",
			Router: `/info`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["poetry/controllers:UserController"] = append(beego.GlobalControllerRouter["poetry/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
