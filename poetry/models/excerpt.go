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

func GetTodayRecommand(id int) (rec *Excerpt, hasNew bool) {
	rec = new(Excerpt)
	o := orm.NewOrm()
	err := o.QueryTable("Excerpt").OrderBy("-id").Limit(1).One(rec)
	if err != nil {
		hasNew = false
		return
	}
	if id == rec.Id {
		hasNew = false
	} else {
		hasNew = true
		o.LoadRelated(rec, "Poem")
		if rec.Poem != nil {
			rec.Poem.TextCn = ""
			o.LoadRelated(rec.Poem, "Poet")
			if rec.Poem.Poet != nil {
				rec.Poem.Poet.Desc = ""
			}
		}
	}
	return
}
