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
    beego.Debug(page)
    var list []*models.Feed
    var err error
    if fid > 0 {
        list, err = models.GetFeedsAfter(fid)
    } else if fid < 0 {
        list = []*models.Feed { }
    } else {
        list, err = models.GetFeeds(page)
    }
	if err != nil {
		c.ReplyErr(err)
		return
	} else {
		c.ReplySucc(list)
	}
}