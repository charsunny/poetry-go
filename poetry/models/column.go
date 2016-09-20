package models

import (
	"errors"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Column))
	beego.Debug("init model poet")
}

type Column struct {
	Id           int
	Title        string
	Desc         string
	Image        string
	Type         int        // 0 表示诗歌 1 表示诗人
	Delete       bool       //是否标记为删除
	Default      bool       // 是否是默认专辑，默认不可以删除
	User         *User      `orm:"rel(fk)"` //创建人
	Poets        []*Poet    `orm:"rel(m2m)"`
	Poems        []*Poem    `orm:"rel(m2m)"`
	IsFav        bool       `orm:"-"`
	Count        int        // 作品数量
	LikeCount    int        // 收藏人数
	CommentCount int        // 评论人数
	Comments     []*Comment `orm:"reverse(many)" json:"-"`
	LikeUsers    []*User    `orm:"reverse(many)" json:"-"`
	Meesages     []*Message `orm:"reverse(many)" json:"-"`
}

func GetColumn(id int) (column *Column, err error) {
	column = new(Column)
	column.Id = id
	err = orm.NewOrm().Read(column)
	return
}

func GetColumnDetail(id int, page int) (column *Column, err error) {
	column, err = GetColumn(id)
	if err != nil {
		return
	}
	o := orm.NewOrm()
	if column.Type == 0 {
		o.LoadRelated(column, "Poems")
		if column.Poems != nil {
			ct := 20
			if page*20+20 > len(column.Poems) {
				ct = len(column.Poems) - page*20
			}
			if ct < 0 {
				ct = 0
			}
			column.Poems = column.Poems[page*20 : ct]
		}
	} else {
		o.LoadRelated(column, "Poets")
		if column.Poets != nil {
			ct := 20
			if page*20+20 > len(column.Poets) {
				ct = len(column.Poets) - page*20
			}
			if ct < 0 {
				ct = 0
			}
			column.Poets = column.Poets[page*20 : ct]
		}
	}
	o.LoadRelated(column, "User")
	return
}

func GetColumnComments(page int, id int) (list []*Comment, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("Comment").Filter("column_id", id).Limit(20).Offset(20 * page).All(&list)
	return
}

func GetColumnFavUsers(page int, id int) (list []*User, err error) {
	col, err := GetColumn(id)
	o := orm.NewOrm()
	o.LoadRelated(col, "LikeUsers")
	if col.LikeUsers != nil && len(col.LikeUsers) > page*20 {
		if len(col.LikeUsers) > (page+1)*20 {
			list = col.LikeUsers[page*20 : (page+1)*20]
		} else {
			list = col.LikeUsers[page*20:]
		}
	}
	return
}

func ColumnUpdatePoem(col *Column, poem *Poem) (add bool, err error) {
	if col.Type != 0 {
		err = errors.New("只能是诗词的专辑才可以添加诗词")
		return
	}
	o := orm.NewOrm()
	exist := o.QueryM2M(col, "Poems").Exist(poem)
	if exist {
		o.QueryM2M(col, "Poems").Remove(poem)
		add = false
		col.Count = col.Count - 1
	} else {
		o.QueryM2M(col, "Poems").Add(poem)
		add = true
		col.Count = col.Count + 1
	}
	_, err = o.Update(col)
	return
}

func ColumnUpdatePoet(col *Column, poet *Poet) (add bool, err error) {
	if col.Type != 1 {
		err = errors.New("只能是诗人的专辑才可以添加诗人")
		return
	}
	o := orm.NewOrm()
	exsit := o.QueryM2M(col, "Poets").Exist(poet)
	if exsit {
		o.QueryM2M(col, "Poets").Remove(poet)
		add = false
		col.Count = col.Count - 1
	} else {
		o.QueryM2M(col, "Poets").Add(poet)
		add = true
		col.Count = col.Count + 1
	}
	_, err = o.Update(col)
	return
}

func ColumnAddComment(col *Column, comment *Comment) (err error) {
	o := orm.NewOrm()
	if comment.Id == 0 {
		comment.Column = col
		_, err = o.Insert(comment)
	}
	col.CommentCount = col.CommentCount + 1
	_, err = o.Update(col)
	return
}

func ColumnUpdateLikeUser(col *Column, user *User) (add bool, err error) {
	o := orm.NewOrm()
	exsit := o.QueryM2M(col, "LikeUsers").Exist(user)
	if exsit {
		o.QueryM2M(col, "LikeUsers").Remove(user)
		add = false
		col.LikeCount = col.LikeCount - 1
	} else {
		o.QueryM2M(col, "LikeUsers").Add(user)
		add = true
		col.LikeCount = col.LikeCount + 1
	}
	_, err = o.Update(col)
	return
}

func AddColumn(col *Column) (column *Column, err error) {
	_, err = orm.NewOrm().Insert(col)
	column = col
	return
}

func DeleteColumn(id, uid int) (err error) {
	col, err := GetColumn(id)
	if err == nil {
		if col.User.Id != uid {
			err = errors.New("不能删除不是自己的专辑")
		} else if col.Default == false {
			err = errors.New("默认专辑无法删除")
		} else {
			col.Delete = true
		}
	}
	orm.NewOrm().Update(col, "Delete")
	return
}
