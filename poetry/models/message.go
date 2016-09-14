package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	beego.Debug("init model poet")
	orm.RegisterModel(new(Message))
}

type Message struct {
	Id         int
	Time       string   // 消息时间
	Type       int      // 1.评论 2. 动态 3. 专辑
	Action     int      // 1. 评论 2. 收藏  3. 喜欢/点赞
	Delete     bool     // 是否已经删除
	User       *User    `orm:"rel(fk)"`                         // 消息的用户
	RelUser    *User    `orm:"rel(one)"`                        // 消息的触发者，一个消息只有一个触发者
	RelComment *Comment `orm:"rel(one);null" json:";omitempty"` // 相关联的评论，当action是1的时候会有
	Comment    *Comment `orm:"rel(fk);null" json:";omitempty"`  // 原始的评论
	Feed       *Feed    `orm:"rel(fk);null" json:";omitempty"`  // 原始的动态
	Column     *Column  `orm:"rel(fk);null" json:";omitempty"`  // 原始的专辑
}

// refresh表示向上还是向下 mid为0表示没有进行查询过
func GetMessageList(uid int, mid int, refresh bool) (list []*Message, err error) {
	o := orm.NewOrm()
	if mid == 0 {
		_, err = o.QueryTable("Message").Filter("user_id", uid).Filter("Delete", false).OrderBy("-Id").Limit(20).All(&list)
	} else {
		if refresh {
			_, err = o.QueryTable("Message").Filter("user_id", uid).Filter("Delete", false).Filter("Id__gt", mid).Limit(20).All(&list)
		} else {
			_, err = o.QueryTable("Message").Filter("user_id", uid).Filter("Delete", false).Filter("Id__lt", mid).Limit(20).All(&list)
		}
	}
	return
}

func DeleteMessage(uid, mid int) (err error) {
	_, err = orm.NewOrm().QueryTable("Message").Filter("user_id", uid).Filter("id", mid).Update(orm.Params{
		"delete": true,
	})
	return
}
