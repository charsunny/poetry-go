package controllers

import (
	"poetry/models"
	"strconv"

	"github.com/astaxie/beego"
	//"lib/ysqi/tokenauth"
	"lib/ysqi/tokenauth2beego/o2o"
)

// Operations about Users
type BaseController struct {
	beego.Controller
	userID int
}

func (b *BaseController) ReplySucc(data interface{}) {
	b.Data["json"] = map[string]interface{}{
		"errcode": 0,
		"data":    data,
	}
	b.ServeJSON()
}

func (b *BaseController) ReplyErr(err error) {
	berr, e := err.(models.BackendError)
	if !e {
		b.Data["json"] = map[string]interface{}{
			"errcode": 1001,
			"errmsg":  err.Error(),
		}
	} else {
		b.Data["json"] = map[string]interface{}{
			"errcode": berr.Code,
			"errmsg":  err.Error(),
		}
	}
	b.ServeJSON()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /login [post]
func (u *BaseController) Login() {
	var user models.User
	user.UserId = u.GetString("userid")
	beego.Debug(u.GetString("userid"))
	if len(user.UserId) == 0 {
		u.ReplyErr(models.BackendError{Code: 1002, Msg: "第三方id为空"})
		return
	}
	user.LoginType, _ = u.GetInt("type")
	user.Avatar = u.GetString("avatar")
	user.Gender, _ = u.GetInt("gender")
	user.Nick = u.GetString("nick")
	models.AddUser(&user)
	token, err := o2o.Auth.NewSingleToken(strconv.Itoa(user.Id), u.Ctx.ResponseWriter)
	if err != nil {
		u.ReplyErr(err)
		return
	}
	u.ReplySucc(map[string]interface{}{
		"token": token.Value,
		"user":  user,
	})
}

// @Title Upload image to server
// @Description get user by uid
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /pic [post]
func (c *BaseController) UploadPic() {
	f, h, err := c.GetFile("image")
	if err == nil {
		path := "./static/pic/" + h.Filename
		f.Close()
		err = c.SaveToFile("image", path)
	}
	path := beego.AppConfig.String("domain") + "static/pic/" + h.Filename
	if err == nil {
		c.ReplySucc(path)
	} else {
		c.ReplyErr(err)
	}
}
