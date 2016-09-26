package models

import (
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	beego.Debug("init model poet")
}

type Feed struct {
	Id           int
	User         *User `orm:"rel(fk)"`
	Time         string
	Content      string
	Picture      string
	Poem         *Poem `orm:"rel(fk)"`
	LikeCount    int
	CommentCount int
	IsFav        bool       `orm:"-"`
	LikeUsers    []*User    `orm:"reverse(many)" json:"-"`
	Comments     []*Comment `orm:"reverse(many)" json:"-"`
	Messages     []*Message `orm:"reverse(many)" json:"-"`
}

func GetFeed(fid int) (p *Feed, err error) {
	p = new(Feed)
	p.Id = fid
	err = orm.NewOrm().Read(p)
	return
}

func AddFeed(uid, pid int, content string, image string) (feed *Feed, err error) {
	feed = new(Feed)
	feed.User = &User{Id: uid}
	feed.Time = time.Now().Format("2006-01-02 15:04:05")
	feed.Content = content
	feed.Picture = image
	feed.Poem = &Poem{Id: pid}
	_, err = orm.NewOrm().Insert(feed)
	return
}

func GetFeeds(uid, page int) (list []*Feed, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("Feed").OrderBy("-Time").Limit(20).Offset(page * 20).All(&list)
	for _, f := range list {
		o.LoadRelated(f, "Poem")
		if f.Poem != nil && len(f.Poem.TextCn) > 200 {
			f.Poem.TextCn = f.Poem.TextCn[0:200] + "..."
			strings.Replace(f.Poem.TextCn, "\r\n", "", 0)
		}
		o.LoadRelated(f, "User")
		if uid > 0 {
			f.IsFav = o.QueryM2M(f, "LikeUsers").Exist(&User{Id: uid})
		}
	}
	return
}

func GetFeedComments(page int, id int, uid int) (list []*Comment, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("Comment").Filter("feed_id", id).OrderBy("-Id").Limit(20).Offset(20 * page).All(&list)
	GetCommentsDetail(uid, list)
	return
}

func GetFeedsAfter(fid int) (list []*Feed, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("Feed").OrderBy("-Time").Filter("id__gt", fid).All(&list)
	for _, f := range list {
		o.LoadRelated(f, "Poem")
		if f.Poem != nil && len(f.Poem.TextCn) > 200 {
			f.Poem.TextCn = f.Poem.TextCn[0:200] + "..."
			strings.Replace(f.Poem.TextCn, "\r\n", "", 0)
		}
		o.LoadRelated(f, "User")
	}
	return
}

func GetUserFeeds(uid int, page int) (list []*Feed, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("Feed").Filter("user_id", uid).OrderBy("-Time").Limit(20).Offset(page * 20).All(&list)
	for _, f := range list {
		o.LoadRelated(f, "Poem")
		if f.Poem != nil && len(f.Poem.TextCn) > 200 {
			f.Poem.TextCn = f.Poem.TextCn[0:200] + "..."
			strings.Replace(f.Poem.TextCn, "\r\n", "", 0)
		}
		o.LoadRelated(f, "User")
	}
	return
}

func GetUserFeedsAfter(uid int, fid int) (list []*Feed, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("Feed").Filter("user_id", uid).OrderBy("-Time").Filter("id__gt", fid).All(&list)
	for _, f := range list {
		o.LoadRelated(f, "Poem")
		if f.Poem != nil && len(f.Poem.TextCn) > 200 {
			f.Poem.TextCn = f.Poem.TextCn[0:200] + "..."
			strings.Replace(f.Poem.TextCn, "\r\n", "", 0)
		}
		o.LoadRelated(f, "User")
	}
	return
}

func LikeFeed(cid int, uid int) (feed *Feed, err error) {
	o := orm.NewOrm()
	feed = new(Feed)
	feed.Id = cid
	if err = o.Read(feed); err != nil {
		return
	}
	exist := o.QueryM2M(feed, "LikeUsers").Exist(&User{Id: uid})
	if !exist {
		o.QueryM2M(feed, "LikeUsers").Add(&User{Id: uid})
		feed.LikeCount = feed.LikeCount + 1
		o.Update(feed, "LikeCount")
	}
	return
}

func DislikeFeed(cid int, uid int) (feed *Feed, err error) {
	o := orm.NewOrm()
	feed = new(Feed)
	feed.Id = cid
	if err = o.Read(feed); err != nil {
		return
	}
	exist := o.QueryM2M(feed, "LikeUsers").Exist(&User{Id: uid})
	if exist {
		o.QueryM2M(feed, "LikeUsers").Remove(&User{Id: uid})
		feed.LikeCount = feed.LikeCount - 1
		o.Update(feed, "LikeCount")
	}
	return
}
