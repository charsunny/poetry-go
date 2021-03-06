package models

import (
	"errors"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Poem), new(PoemFormat))
	beego.Debug("init model poem")
}

type Poem struct {
	Id           int
	NameCn       string `json:"Name"`
	Poet         *Poet  `orm:"rel(fk)"`
	TextCn       string `json:"Text"`
	FormatId     int    `json:"-"`
	Format       string `orm:"-"`
	IsFav        bool   `orm:"-"`
	LikeCount    int
	CommentCount int
	Comments     []*Comment   `orm:"reverse(many)" json:"-"`
	Recommands   []*Recommand `orm:"reverse(many)" json:"-"`
	LikeUsers    []*User      `orm:"reverse(many)" json:"-"`
	Columns      []*Column    `orm:"reverse(many)" json:",omitempty"` // 收藏的相关专辑
	Excerpts     []*Excerpt   `orm:"reverse(many)" json:"-"`          // 关联的摘抄
}

type PoemFormat struct {
	Id     int
	NameCn string `json:"Name"`
	Desc   string `orm:"column(description_cn)"`
}

func GetPoem(pid int) (p *Poem, err error) {
	p = new(Poem)
	p.Id = pid
	err = orm.NewOrm().Read(p)
	return
}

func GetPoemByRow(id int) (p *Poem, err error) {
	p = new(Poem)
	o := orm.NewOrm()
	err = o.QueryTable("Poem").Limit(1).Offset(id).One(p)
	o.LoadRelated(p, "Poet")
	return
}

func GetPoemFormat(id int) (p *PoemFormat, err error) {
	p = new(PoemFormat)
	p.Id = id
	err = orm.NewOrm().Read(p)
	return
}

func GetPoemFormatByRow(id int) (p *PoemFormat, err error) {
	p = new(PoemFormat)
	err = orm.NewOrm().QueryTable("PoemFormat").Limit(1).Offset(id).One(p)
	return
}

func GetPoemDetail(pid, uid int) (p *Poem, err error) {
	p, err = GetPoem(pid)

	o := orm.NewOrm()
	o.LoadRelated(p, "Poet")

	if uid > 0 {
		p.IsFav = o.QueryM2M(p, "LikeUsers").Exist(&User{Id: uid})
	}

	// if p.Poet != nil {
	//     p.Poet.Desc = ""
	// }
	return
}

func FavPoem(pid, uid int) (fav bool, err error) {
	p, err := GetPoem(pid)
	o := orm.NewOrm()
	if uid > 0 {
		u := &User{Id: uid}
		query := o.QueryM2M(u, "LikePoems")
		isFav := query.Exist(p)
		if isFav {
			if _, err = query.Remove(p); err == nil {
				p.LikeCount--
				fav = false
			}
		} else {
			if _, err = query.Add(p); err == nil {
				p.LikeCount++
				fav = true
			}
		}
		o.Update(p)
	} else {
		err = errors.New("用户尚未登录")
	}
	return
}

func GetPoemComments(pid, page int, uid int) (list []*Comment, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("Comment").Filter("poem_id", pid).OrderBy("-Id").Limit(20).Offset(page * 20).All(&list)
	GetCommentsDetail(uid, list)
	return
}
