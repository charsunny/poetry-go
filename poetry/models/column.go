package models

import (
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego"
)

func init()  {
    orm.RegisterModel(new(Column))
    beego.Debug("init model poet")
}

type Column struct {
    Id int 
    Title string   
    Desc   string
    Type int // 0 表示诗歌 1 表示诗人
    User *User          `orm:"rel(fk)"`    //创建人
    Poets []*Poet      `orm:"rel(m2m)"`
    Poems []*Poem      `orm:"rel(m2m)"`
    IsFav bool    `orm:"-"`
    LikeCount int       // 收藏人数
    CommentCount int    // 评论人数
    Comments []*Comment     `orm:"reverse(many)" json:"-"`
    LikeUsers []*User `orm:"reverse(many)" json:"-"`
}