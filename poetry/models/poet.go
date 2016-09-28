package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Period), new(Poet))
	beego.Debug("init model poet")
}

type Period struct {
	Id     int
	NameCn string
}

type Poet struct {
	Id        int
	NameCn    string `json:"Name"`
	Avatar    string
	LikeCount int
	Period    string    `orm:"-"`
	PeriodId  int       `json:"-"`
	Desc      string    `orm:"column(description_cn)"`
	Poems     []*Poem   `orm:"reverse(many)" json:",omitempty"`
	LikeUsers []*User   `orm:"reverse(many)" json:",omitempty"`
	Columns   []*Column `orm:"reverse(many)" json:",omitempty"`
}

func GetPoet(pid int) (p *Poet, err error) {
	p = new(Poet)
	p.Id = pid
	err = orm.NewOrm().Read(p)
	return
}

func GetPoetByName(name string) (p *Poet, err error) {
	p = new(Poet)
	err = orm.NewOrm().QueryTable("Poet").Filter("NameCn", name).One(p)
	return
}

func GetPoetByRow(id int) (p *Poet, err error) {
	p = new(Poet)
	err = orm.NewOrm().QueryTable("Poet").Limit(1).Offset(id).One(p)
	return
}
