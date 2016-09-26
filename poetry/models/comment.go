package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	beego.Debug("init model poet")
}

type Comment struct {
	Id        int
	Comment   *Comment   `orm:"rel(fk);null"`           // 评论别的评论
	RComments []*Comment `orm:"reverse(many)" json:"-"` // 这个评论被别的评论评论
	User      *User      `orm:"rel(fk)"`
	Time      string
	Content   string
	Picture   string
	LikeCount int
	IsFav     bool  `orm:"-"`
	Feed      *Feed `orm:"rel(fk);null"`
	// 对应message 不做展示
	Column     *Column    `orm:"rel(fk);null"`
	LikeUsers  []*User    `orm:"reverse(many)" json:"-"`
	Poem       *Poem      `orm:"rel(fk);null", json:"-"` // 评论别的评论
	RelMessage *Message   `orm:"reverse(one)" json:"-"`
	Messages   []*Message `orm:"reverse(many)" json:"-"`
}

func AddComment(tp, id, uid, cid int, content string) (comment *Comment, err error) {
	o := orm.NewOrm()
	user, err := GetUser(uid)
	if err != nil {
		beego.Debug(err)
		return
	}
	comment = new(Comment)
	if cid > 0 {
		comment.Comment = &Comment{Id: cid}
	}
	comment.Content = content
	comment.User = user
	comment.Time = time.Now().Format("01-02 15:04:05")
	if tp == 1 {
		poem, err := GetPoem(id)
		if err != nil {
			beego.Debug(err)
			return comment, err
		}
		comment.Poem = poem
		_, err = o.Insert(comment)
		if err != nil {
			beego.Debug(err)
			return comment, err
		}
		poem.CommentCount = poem.CommentCount + 1
		_, err = o.Update(poem)
	} else if tp == 2 {
		feed, err := GetFeed(id)
		if err != nil {
			return comment, err
		}
		comment.Feed = feed
		_, err = o.Insert(comment)
		if err != nil {
			return comment, err
		}
		feed.CommentCount = feed.CommentCount + 1
		_, err = o.Update(feed)
	} else {
		col, err := GetColumn(id)
		if err != nil {
			return comment, err
		}
		comment.Column = col
		_, err = o.Insert(comment)
		if err != nil {
			return comment, err
		}
		col.CommentCount = col.CommentCount + 1
		_, err = o.Update(col)
	}
	o.LoadRelated(comment, "Comment")
	return
}

func GetCommentsDetail(uid int, list []*Comment) {
	o := orm.NewOrm()
	for _, c := range list {
		o.LoadRelated(c, "Comment")
		if c.Comment != nil {
			o.LoadRelated(c.Comment, "User")
		}
		o.LoadRelated(c, "User")
		if uid > 0 {
			c.IsFav = o.QueryM2M(c, "LikeUsers").Exist(&User{Id: uid})
		}
	}
}

func LikeComment(cid int, uid int) (comment *Comment, err error) {
	o := orm.NewOrm()
	comment = new(Comment)
	comment.Id = cid
	if err = o.Read(comment); err != nil {
		return
	}
	exist := o.QueryM2M(comment, "LikeUsers").Exist(&User{Id: uid})
	if !exist {
		o.QueryM2M(comment, "LikeUsers").Add(&User{Id: uid})
		comment.LikeCount = comment.LikeCount + 1
		o.Update(comment, "LikeCount")
	}
	return
}

func DislikeComment(cid int, uid int) (comment *Comment, err error) {
	o := orm.NewOrm()
	comment = new(Comment)
	comment.Id = cid
	if err = o.Read(comment); err != nil {
		return
	}
	exist := o.QueryM2M(comment, "LikeUsers").Exist(&User{Id: uid})
	if exist {
		o.QueryM2M(comment, "LikeUsers").Remove(&User{Id: uid})
		comment.LikeCount = comment.LikeCount - 1
		o.Update(comment, "LikeCount")
	}
	return
}
