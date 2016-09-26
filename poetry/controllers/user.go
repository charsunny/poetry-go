package controllers

import (
	"poetry/models"
	"strconv"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	BaseController
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	user.Avatar = u.GetString("avatar")
	user.UserId = u.GetString("uid")
	user.Gender, _ = u.GetInt("gender")
	user.Nick = u.GetString("nick")
	models.AddUser(&user)
	u.ReplySucc(user)
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		id, _ := strconv.Atoi(uid)
		models.GetUser(id)
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /addfeed [post]
func (u *UserController) AddFeed() {
	uid, err := u.GetInt("uid")
	if err != nil {
		u.ReplyErr(err)
		return
	}
	pid, err := u.GetInt("pid")
	if err != nil {
		u.ReplyErr(err)
		return
	}
	image := u.GetString("image")
	content := u.GetString("content")
	_, err = models.AddFeed(uid, pid, content, image)
	if err != nil {
		u.ReplyErr(err)
		return
	} else {
		u.ReplySucc("OK")
	}
}

// @Title get user feeds
// @Description get user feeds
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /feeds [get]
func (u *UserController) GetFeeds() {
	uid, err := u.GetInt("uid")
	id, _ := u.GetInt("id")
	if id > 0 {
		uid = id
	}
	beego.Debug(uid)
	if err != nil {
		u.ReplyErr(err)
		return
	}
	page, _ := u.GetInt("page")
	fid, _ := u.GetInt("fid")
	beego.Debug(page)
	var list []*models.Feed
	if fid > 0 {
		list, err = models.GetFeedsAfter(fid)
	} else if fid < 0 {
		list = []*models.Feed{}
	} else {
		list, err = models.GetUserFeeds(uid, page)
	}
	if err != nil {
		u.ReplyErr(err)
		return
	} else {
		u.ReplySucc(list)
	}
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /info [get]
func (u *UserController) GetInfo() {
	uid, _ := u.GetInt("uid")
	beego.Debug(uid)
	user, err := models.GetUser(uid)
	if err != nil {
		u.ReplyErr(err)
	} else {
		u.ReplySucc(user)
	}
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

// @Title  add comment
// @Description update the user
// @Param	page		query 	int	true		"The page you want to get, default is 0"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /addcomment [post]
func (c *UserController) AddComment() {
	cid, _ := c.GetInt("cid") // 评论的id
	id, _ := c.GetInt("id")   // 专辑的id
	uid, _ := c.GetInt("uid")
	content := c.GetString("content")
	tp, _ := c.GetInt("type")
	cm, err := models.AddComment(tp, id, uid, cid, content)
	if err != nil {
		beego.Debug(err)
		c.ReplyErr(err)
	} else {
		c.ReplySucc(cm)
	}
}

// @Title  like comment
// @Description update the user
// @Param	page		query 	int	true		"The page you want to get, default is 0"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /likecomment [post]
func (c *UserController) LikeComment() {
	cid, _ := c.GetInt("cid") // 评论的id
	uid, _ := c.GetInt("uid")
	cm, err := models.LikeComment(cid, uid)
	if err != nil {
		beego.Debug(err)
		c.ReplyErr(err)
	} else {
		c.ReplySucc(cm)
	}
}

// @Title  like comment
// @Description update the user
// @Param	page		query 	int	true		"The page you want to get, default is 0"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /unlikecomment [post]
func (c *UserController) DislikeComment() {
	cid, _ := c.GetInt("cid") // 评论的id
	uid, _ := c.GetInt("uid")
	cm, err := models.DislikeComment(cid, uid)
	if err != nil {
		beego.Debug(err)
		c.ReplyErr(err)
	} else {
		c.ReplySucc(cm)
	}
}

// @Title  like comment
// @Description update the user
// @Param	page		query 	int	true		"The page you want to get, default is 0"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /likefeed [post]
func (c *UserController) LikeFeed() {
	cid, _ := c.GetInt("cid") // 评论的id
	uid, _ := c.GetInt("uid")
	cm, err := models.LikeFeed(cid, uid)
	if err != nil {
		beego.Debug(err)
		c.ReplyErr(err)
	} else {
		c.ReplySucc(cm)
	}
}

// @Title  like comment
// @Description update the user
// @Param	page		query 	int	true		"The page you want to get, default is 0"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /unlikefeed [post]
func (c *UserController) DislikeFeed() {
	cid, _ := c.GetInt("cid") // 评论的id
	uid, _ := c.GetInt("uid")
	cm, err := models.DislikeFeed(cid, uid)
	if err != nil {
		beego.Debug(err)
		c.ReplyErr(err)
	} else {
		c.ReplySucc(cm)
	}
}
