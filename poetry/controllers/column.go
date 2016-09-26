package controllers

import (
	"errors"
	"poetry/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// Operations about Users
type ColumnController struct {
	BaseController
}

// @Title Get Recommand Info
// @Description get recommd by rid
// @Param	rid	 query 	int	false		"The key for Recommand, = 0 get last"
// @Success 200 {object} models.Recommand
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ColumnController) GetColumn() {
	id := c.GetString(":id")
	page, _ := c.GetInt("page")
	if id != "" {
		cid, _ := strconv.Atoi(id)
		col, err := models.GetColumnDetail(cid, page)
		if err != nil {
			c.ReplyErr(err)
		} else {
			c.ReplySucc(col)
		}
	} else {
		err := errors.New("参数错误")
		c.ReplyErr(err)
	}
}

// @Title Get Recommand Info
// @Description get recommd by rid
// @Param	rid	 query 	int	false		"The key for Recommand, = 0 get last"
// @Success 200 {object} models.Recommand
// @Failure 403 :uid is empty
// @router /user [get]
func (c *ColumnController) GetUserColumns() {
	uid, _ := c.GetInt("uid")
	list, _, err := models.GetUserColumns(uid)
	if err != nil {
		c.ReplyErr(err)
	} else {
		if len(list) == 0 {
			user, err := models.GetUser(uid)
			if err != nil {
				c.ReplyErr(err)
				return
			}
			// do create default list action
			col := new(models.Column)
			col.Title = "喜欢的诗词"
			col.Desc = "喜欢的诗词"
			col.Type = 0
			col.Default = true
			col.User = user
			col, err = models.AddColumn(col)
			if err == nil {
				list = append(list, col)
			}
			beego.Debug(err)
			col = new(models.Column)
			col.Title = "喜欢的诗人"
			col.Desc = "喜欢的诗人"
			col.Type = 1
			col.Default = true
			col.User = user
			col, err = models.AddColumn(col)
			if err == nil {
				list = append(list, col)
			}
			beego.Debug(err)
		}
		c.ReplySucc(list)
	}
}

// @Title Get Recommand Info
// @Description get recommd by rid
// @Param	rid	 query 	int	false		"The key for Recommand, = 0 get last"
// @Success 200 {object} models.Recommand
// @Failure 403 :uid is empty
// @router /userfav [get]
func (c *ColumnController) GetUserFavColumns() {
	uid, _ := c.GetInt("uid")
	_, likelist, err := models.GetUserColumns(uid)
	if err != nil {
		c.ReplyErr(err)
	} else {
		c.ReplySucc(likelist)
	}
}

// @Title Get Recommand Info
// @Description get recommd by rid
// @Param	rid	 query 	int	false		"The key for Recommand, = 0 get last"
// @Success 200 {object} models.Recommand
// @Failure 403 :uid is empty
// @router /updateitem [post]
func (c *ColumnController) UpdateColumnItem() {
	pid, _ := c.GetInt("pid")
	cid, _ := c.GetInt("cid")
	uid, _ := c.GetInt("uid")
	col, err := models.GetColumn(cid)
	if err != nil {
		c.ReplyErr(err)
		return
	}
	if col.User.Id != uid {
		c.ReplyErr(err)
		return
	}
	add := false
	if col.Type == 1 {
		poet, err := models.GetPoet(pid)
		if err != nil {
			c.ReplyErr(err)
			return
		}
		add, err = models.ColumnUpdatePoet(col, poet)
	} else {
		poem, err := models.GetPoem(pid)
		if err != nil {
			c.ReplyErr(err)
			return
		}
		add, err = models.ColumnUpdatePoem(col, poem)
	}
	if err != nil {
		c.ReplyErr(err)
		return
	}
	if add {
		c.ReplySucc("Update")
	} else {
		c.ReplySucc("Delete")
	}
}

// @Title  Get Recommand List
// @Description update the user
// @Param	page		query 	int	true		"The page you want to get, default is 0"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /comments [get]
func (c *ColumnController) GetColumnComments() {
	page, _ := c.GetInt("page")
	cid, _ := c.GetInt("cid")
	uid, _ := c.GetInt("uid")
	list, err := models.GetColumnComments(page, cid, uid)
	if err != nil {
		c.ReplyErr(err)
	} else {
		c.ReplySucc(list)
	}
}

// @Title  Get Recommand List
// @Description update the user
// @Param	page		query 	int	true		"The page you want to get, default is 0"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /favusers [get]
func (c *ColumnController) GetColumnUsers() {
	page, _ := c.GetInt("page")
	cid, _ := c.GetInt("cid")
	list, err := models.GetColumnFavUsers(page, cid)
	if err != nil {
		c.ReplyErr(err)
	} else {
		c.ReplySucc(list)
	}
}

// @Title  Get Recommand List
// @Description update the user
// @Param	page		query 	int	true		"The page you want to get, default is 0"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /updatefav [post]
func (c *ColumnController) UpdateFavColumn() {
	cid, _ := c.GetInt("cid")
	uid, _ := c.GetInt("uid")
	col, err := models.GetColumn(cid)
	if err != nil {
		c.ReplyErr(err)
		return
	}
	if col.User.Id == uid {
		c.ReplyErr(err)
		return
	}
	user, err := models.GetUser(uid)
	if err != nil {
		c.ReplyErr(err)
		return
	}
	add, err := models.ColumnUpdateLikeUser(col, user)
	if err != nil {
		c.ReplyErr(err)
		return
	}
	if add {
		c.ReplySucc("Update")
	} else {
		c.ReplySucc("Delete")
	}
}

// @Title  Get Recommand List
// @Description update the user
// @Param	page		query 	int	true		"The page you want to get, default is 0"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /create [post]
func (c *ColumnController) CreateColumn() {
	uid, err := c.GetInt("uid")
	if err != nil {
		c.ReplyErr(err)
		return
	}
	user, err := models.GetUser(uid)
	if err != nil {
		c.ReplyErr(err)
		return
	}
	name := c.GetString("name")
	desc := c.GetString("desc")
	image := c.GetString("image")
	tp, _ := c.GetInt("type")
	contents := c.GetString("contents")
	col := new(models.Column)
	col.Title = name
	col.Desc = desc
	col.Image = image
	col.Type = tp
	col.User = user
	col, err = models.AddColumn(col)
	if err != nil {
		c.ReplyErr(err)
	} else {
		for _, c := range strings.Split(contents, "|") {
			c = strings.Trim(c, " ")
			id, _ := strconv.Atoi(c)
			if id > 0 {
				if tp == 0 {
					models.ColumnUpdatePoem(col, &models.Poem{Id: id})
				} else {
					models.ColumnUpdatePoet(col, &models.Poet{Id: id})
				}
			}
		}
		c.ReplySucc(col)
	}
}

// @Title  Get Recommand List
// @Description update the user
// @Param	page		query 	int	true		"The page you want to get, default is 0"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /delete [post]
func (c *ColumnController) DeleteColumn() {
	id, _ := c.GetInt("id")
	uid, _ := c.GetInt("uid")
	err := models.DeleteColumn(id, uid)
	if err != nil {
		c.ReplyErr(err)
	} else {
		c.ReplySucc("OK")
	}
}
