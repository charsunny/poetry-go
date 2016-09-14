package controllers

import (
	"errors"
	"poetry/models"
	"strconv"
)

// Operations about Users
type MessageController struct {
	BaseController
}

// @Title Get Recommand Info
// @Description get recommd by rid
// @Param	rid	 query 	int	false		"The key for Recommand, = 0 get last"
// @Success 200 {object} models.Recommand
// @Failure 403 :id is empty
// @router /:id [get]
func (c *MessageController) GetMessage() {
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
// @router /list [get]
func (c *MessageController) GetMessageList() {
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

// @Title  delte Mesage List
// @Description update the user
// @Param	page		query 	int	true		"The page you want to get, default is 0"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /delete [post]
func (c *MessageController) DeleteMessage() {
	id, _ := c.GetInt("id")
	uid, _ := c.GetInt("uid")
	err := models.DeleteColumn(id, uid)
	if err != nil {
		c.ReplyErr(err)
	} else {
		c.ReplySucc("OK")
	}
}
