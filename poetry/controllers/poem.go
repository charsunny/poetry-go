package controllers

import (
	"errors"
	"poetry/models"
	"strconv"

	_ "github.com/astaxie/beego"
)

// Operations about Users
type PoemController struct {
	BaseController
}

// @Title Get
// @Description get poem by pid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :pid is empty
// @router /:pid [get]
func (c *PoemController) Get() {
	pid := c.Ctx.Input.Param(":pid")
	haslocal, _ := c.GetBool("local")
	uid, _ := c.GetInt("uid")
	if pid != "" {
		id, _ := strconv.Atoi(pid)
		p, err := models.GetPoemDetail(id, uid)
		if haslocal {
			p.Poet = nil
			p.TextCn = ""
		}
		if err != nil {
			c.ReplyErr(err)
		} else {
			c.ReplySucc(p)
		}
	} else {
		c.ReplyErr(errors.New("请求参数错误"))
	}
}

// @Title Get
// @Description get poem comments by pid and page
// @Param	pid   query int	true		"The key for staticblock"
// @Param	page  query int	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :pid is empty
// @router /comments [get]
func (c *PoemController) GetComments() {
	pid, _ := c.GetInt("pid")
	page, _ := c.GetInt("page")
	uid, _ := c.GetInt("uid")
	list, err := models.GetPoemComments(pid, page, uid)
	if err != nil {
		c.ReplyErr(err)
	} else {
		c.ReplySucc(list)
	}
}

// @Title Post Like
// @Description fav poem by pid
// @Param	pid   query int	true		"The key for staticblock"
// @Param	page  query int	true		"The key for staticblock"
// @Success 200 {string} like state
// @Failure 403 :pid is empty
// @router /like [post]
func (c *PoemController) LikePoem() {
	pid, _ := c.GetInt("pid")
	uid, _ := c.GetInt("uid")
	fav, err := models.FavPoem(pid, uid)
	if err != nil {
		c.ReplyErr(err)
	} else {
		if fav {
			c.ReplySucc("已添加喜欢")
		} else {
			c.ReplySucc("已取消喜欢")
		}
	}
}
