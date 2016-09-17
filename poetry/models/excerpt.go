package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	beego.Debug("init model poet")
	orm.RegisterModel(new(Excerpt))
}

type Excerpt struct {
	Id      int
	Time    string // 时间
	Content string // 内容
	Poem    *Poem  `orm:"rel(fk)"` // 关联诗词
}
