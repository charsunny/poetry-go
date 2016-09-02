package models

import (
    "time"
    "strings"
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego"
)

func init()  {
    beego.Debug("init model poet")
}

type Feed struct {
    Id int 
    User *User `orm:"rel(fk)"`
    Time string 
    Content string 
    Picture string
    Poem *Poem `orm:"rel(fk)"`
    LikeCount int 
    CommentCount int 
    LikeUsers []*User `orm:"reverse(many)" json:"-"`
    Comments []*Comment `orm:"reverse(many)" json:"-"`
    Messages []*Message `orm:"reverse(many)" json:"-"`
}

func AddFeed(uid, pid int, content string, image string) (feed *Feed, err error)  {
    feed = new(Feed)
    feed.User = &User{Id:uid}
    feed.Time = time.Now().Format("2006-01-02 15:04:05")
    feed.Content = content
    feed.Picture = image
    feed.Poem = &Poem{Id:pid}
    _, err = orm.NewOrm().Insert(feed)
    return 
}

func GetFeeds(page int) (list []*Feed, err error) {
    o := orm.NewOrm()
    _, err = o.QueryTable("Feed").OrderBy("-Time").Limit(20).Offset(page*20).All(&list)
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
    _, err = o.QueryTable("Feed").Filter("user_id", uid).OrderBy("-Time").Limit(20).Offset(page*20).All(&list)
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
