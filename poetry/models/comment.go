package models

import (
    _ "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego"
)

func init()  {
    beego.Debug("init model poet")
}

type Comment struct {
    Id int 
    Comment *Comment`orm:"rel(fk)"`           // 评论别的评论
    RComments []*Comment`orm:"reverse(many)" json:"-"` // 这个评论被别的评论评论
    User *User `orm:"rel(fk)"`
    Time string 
    Content string 
    Picture string
    LikeCount int 
    Feed *Feed `orm:"rel(fk)"`
    // 对应message 不做展示
    Poem *Poem `orm:"rel(fk)", json:"-"`           // 评论别的评论
    RelMessage *Message `orm:"reverse(one)" json:"-"`
    Messages []*Message `orm:"reverse(many)" json:"-"`
}