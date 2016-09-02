package models

import (
     "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego"
)

func init()  {
    beego.Debug("init model poet")
    orm.RegisterModel(new(Message))
}

type Message struct {
    Id int 
    Time string                             // 消息时间
    User *User `orm:"rel(fk)"`              // 消息的用户
    RelUser *User `orm:"rel(one)"`          // 谁干的
    RelComment  *Comment `orm:"rel(one)"`   //干的人的评论， 如果没有 就是赞
    Comment *Comment `orm:"rel(fk)"`        // 消息用户的原始评论， 赞了评论， 回复了评论
    Feed *Feed `orm:"rel(fk)"`              // 消息用户的原始状态, 赞了状态， 恢复了状态
}