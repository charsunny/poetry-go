package controllers

import (
	"errors"
	"poetry/models"
	"strconv"
	"strings"
	"time"
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
	if id != "" {
		cid, _ := strconv.Atoi(id)
		col, err := models.GetColumnDetail(cid)
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
			// do create default list action
			col := new(models.Column)
			col.Title = "喜欢的诗词"
			col.Desc = "喜欢的诗词"
			col.Type = 0
			col.Default = true
			col, err := models.AddColumn(col)
			if err != nil {
				list = append(list, col)
			}
			col = new(models.Column)
			col.Title = "喜欢的诗人"
			col.Desc = "喜欢的诗人"
			col.Type = 1
			col.Default = true
			col, err = models.AddColumn(col)
			if err != nil {
				list = append(list, col)
			}
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
// @router /updateitem [get]
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
	list, err := models.GetColumnComments(page, cid)
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
	col, err := models.AddColumn(col)
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

// @Title  add comment
// @Description update the user
// @Param	page		query 	int	true		"The page you want to get, default is 0"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /addcomment [post]
func (c *ColumnController) ColumnAddComment() {
	cid, _ := c.GetInt("cid") // 评论的id
	id, _ := c.GetInt("id")   // 专辑的id
	uid, _ := c.GetInt("uid")
	user, err := models.GetUser(uid)

	content := c.GetString("content")

	col, err := models.GetColumn(id)

	if err != nil {
		c.ReplyErr(err)
	} else {
		comment := new(models.Comment)
		comment.Content = content
		comment.Time = time.Now().Format("2006-01-02 15:04:045")
		if cid > 0 {
			comment.Comment = &models.Comment{Id: cid}
		}
		comment.User = user
		comment.Column = col
		err = models.ColumnAddComment(col, comment)
		if err != nil {
			c.ReplyErr(err)
		} else {
			c.ReplySucc(comment)
		}
	}
}
