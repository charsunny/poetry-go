package models

import (
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego"
)

var (
    RecommandList []*Recommand
)

func init()  {
    orm.RegisterModel(new(Recommand))
    beego.Debug("init model poet")
}

type Recommand struct {
    Id int 
    Title string 
    Desc string 
    Time string
    Poems []*Poem `orm:"rel(m2m)" json:",omitempty"`
}

func GetRecommand(id int) (rec *Recommand, err error) {
    o := orm.NewOrm()
    rec = &Recommand{Id: id}
    err = o.Read(rec)
    if err != nil {
        err = o.QueryTable("Recommand").OrderBy("-Id").Limit(1).One(rec)
    }
    o.LoadRelated(rec, "Poems")
    if rec.Poems != nil {
        var pids []interface{}
        for _, poem := range rec.Poems {
            if len(poem.TextCn) > 60 {
                poem.TextCn = poem.TextCn[0:60]
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
    _, err = o.QueryTable("Recommand").OrderBy("-Id").Limit(20).Offset(page*20).All(&list)
    return 
}