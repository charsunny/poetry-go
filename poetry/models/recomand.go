package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var (
	RecommandList []*Recommand
)

func init() {
	orm.RegisterModel(new(Recommand))
	beego.Debug("init model poet")
}

type Recommand struct {
	Id      int
	Title   string
	Desc    string
	Time    string
	User    *User   `orm:"rel(fk)" json:",omitempty"`
	Poems   []*Poem `orm:"rel(m2m)" json:",omitempty"`
	Publish bool
}

func GetRecommand(id int) (rec *Recommand, err error) {
	o := orm.NewOrm()
	rec = &Recommand{Id: id}
	err = o.Read(rec)
	if err != nil {
		err = o.QueryTable("Recommand").Filter("Publish", true).OrderBy("-Id").Limit(1).One(rec)
	}
	o.LoadRelated(rec, "Poems")
	o.LoadRelated(rec, "User")
	if rec.Poems != nil {
		var pids []interface{}
		for _, poem := range rec.Poems {
			if len(poem.TextCn) > 120 {
				poem.TextCn = poem.TextCn[0:120]
			}
			if poem.Poet != nil {
				pids = append(pids, poem.Poet.Id)
			}
		}
		if len(pids) > 0 {
			var poets []*Poet
			o.QueryTable("Poet").Filter("id__in", pids...).All(&poets)
			for _, p := range poets {
				for _, pm := range rec.Poems {
					if pm.Poet != nil && p.Id == pm.Poet.Id {
						p.Desc = ""
						pm.Poet = p
					}
				}
			}
		}
	}
	return
}

func GetRecommandList(page int) (list []*Recommand, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("Recommand").Filter("Publish", true).OrderBy("-Id").Limit(20).Offset(page * 20).All(&list)
	return
}

func AddRecommand(col *Recommand) (column *Recommand, err error) {
	_, err = orm.NewOrm().Insert(col)
	column = col
	return
}

func RecommandUpdatePoem(col *Recommand, poem *Poem) (add bool, err error) {
	o := orm.NewOrm()
	exist := o.QueryM2M(col, "Poems").Exist(poem)
	if !exist {
		o.QueryM2M(col, "Poems").Add(poem)
		add = true
	}
	_, err = o.Update(col)
	return
}
