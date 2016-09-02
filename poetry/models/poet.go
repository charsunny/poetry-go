package models

import (
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego"
)

func init()  {
    orm.RegisterModel(new(Period), new(Poet))
    beego.Debug("init model poet")
}

type Period struct {
    Id int 
    NameCn string
}

type Poet struct {
    Id int
    NameCn string `json:"Name"`
    Avatar string 
    LikeCount int 
    Period string `orm:"-"` 
    PeriodId int `json:"-"`
    Desc string `orm:"column(description_cn)"`
    Poems []*Poem   `orm:"reverse(many)" json:",omitempty"`
    LikeUsers []*User `orm:"reverse(many)" json:",omitempty"`
}