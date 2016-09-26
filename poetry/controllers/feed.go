package controllers

import (
	"poetry/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type FeedController struct {
	BaseController
}

// @Title get all feeds
// @Description get user feeds
// @Param	page	 path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /list [get]
func (c *FeedController) GetFeeds() {
	page, _ := c.GetInt("page")
	fid, _ := c.GetInt("fid")
	uid, _ := c.GetInt("uid")
	id, _ := c.GetInt("id")
	if id > 0 {
		uid = id
	}
	beego.Debug(page)
	var list []*models.Feed
	var err error
	if fid > 0 {
		list, err = models.GetFeedsAfter(fid)
	} else if fid < 0 {
		list = []*models.Feed{}
	} else {
		list, err = models.GetFeeds(uid, page)
	}
	if err != nil {
		c.ReplyErr(err)
		return
	} else {
		c.ReplySucc(list)
	}
}

// @Title  Get Recommand List
// @Description update the user
// @Param	page		query 	int	true		"The page you want to get, default is 0"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /comments [get]
func (c *FeedController) GetFeedComments() {
	page, _ := c.GetInt("page")
	cid, _ := c.GetInt("fid")
	uid, _ := c.GetInt("uid")
	list, err := models.GetFeedComments(page, cid, uid)
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
func (c *FeedController) GetFeedFavUsers() {
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
// @router /delete [post]
func (c *FeedController) DeleteFeed() {
	id, _ := c.GetInt("id")
	uid, _ := c.GetInt("uid")
	err := models.DeleteColumn(id, uid)
	if err != nil {
		c.ReplyErr(err)
	} else {
		c.ReplySucc("OK")
	}
}
