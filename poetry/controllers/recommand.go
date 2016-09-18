package controllers

import (
	"errors"
	"poetry/models"
)

// Operations about Users
type RecommandController struct {
	BaseController
}

// @Title Get Recommand today info
// @Description get recommd by rid
// @Param	rid	 query 	int	false		"The key for Recommand, = 0 get last"
// @Success 200 {object} models.Recommand
// @Failure 403 :uid is empty
// @router /today [get]
func (u *RecommandController) GetToday() {
	id, _ := u.GetInt("id")
	rec, hasNew := models.GetTodayRecommand(id)
	if !hasNew {
		err := errors.New("没有推荐诗词")
		u.ReplyErr(err)
	} else {
		u.ReplySucc(rec)
	}
}

// @Title Get Recommand Info
// @Description get recommd by rid
// @Param	rid	 query 	int	false		"The key for Recommand, = 0 get last"
// @Success 200 {object} models.Recommand
// @Failure 403 :uid is empty
// @router /info [get]
func (u *RecommandController) Get() {
	rid, _ := u.GetInt("rid")
	rec, err := models.GetRecommand(rid)
	if err != nil {
		u.ReplyErr(err)
	} else {
		u.ReplySucc(rec)
	}
}

// @Title  Get Recommand List
// @Description update the user
// @Param	page		query 	int	true		"The page you want to get, default is 0"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /list [get]
func (u *RecommandController) GetList() {
	page, _ := u.GetInt("page")
	list, err := models.GetRecommandList(page)
	if err != nil {
		u.ReplyErr(err)
	} else {
		u.ReplySucc(list)
	}
}
