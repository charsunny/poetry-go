package models

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/astaxie/beego/orm"
	"github.com/rongcloud/server-sdk-go/RCServerSDK"
)

var (
	UserList map[string]*User
	rcServer *RCServerSDK.RCServer
)

const (
	LoginTypeWeiBo = iota + 1
	LoginTypeQQ
)

func init() {
	UserList = make(map[string]*User)
	orm.RegisterModel(new(User), new(Comment), new(Feed))
	rcServer, _ = RCServerSDK.NewRCServer("8w7jv4qb77vey", "UNIDKJuSiB", "json")
}

type User struct {
	Id            int
	Nick          string
	Avatar        string
	Gender        int
	UserId        string //第三方平台Id
	LoginType     int    //第三方平台类型
	LikeCount     int    //喜欢的诗词数量
	ColumnCount   int    // 专辑数量
	FollowCount   int    // 关注的数量
	FolloweeCount int    // 被关注的数量
	RongUser      string // rongyun用户名
	RongToken     string // rong token
	// Relations
	Feeds       []*Feed    `orm:"reverse(many)"`
	Columns     []*Column  `orm:"reverse(many)"  json:",omitempty"`
	LikeFeeds   []*Feed    `orm:"rel(m2m)"  json:",omitempty"`
	LikePoems   []*Poem    `orm:"rel(m2m)"  json:",omitempty"`
	LikePoets   []*Poet    `orm:"rel(m2m)"  json:",omitempty"`
	LikeColumns []*Column  `orm:"rel(m2m)"  json:",omitempty"`
	Comments    []*Comment `orm:"reverse(many)"`
	Messages    []*Message `orm:"reverse(many)"`

	// dont show
	RelMessage *Message `orm:"reverse(one)" json:"-"`
}

func AddUser(u *User) bool {
	o := orm.NewOrm()
	ct, _ := o.QueryTable("User").Filter("UserId", u.UserId).Count()
	if ct == 0 {
		o.Insert(u)
		return true
	}
	o.QueryTable("User").Filter("UserId", u.UserId).One(u)
	if len(u.RongUser) == 0 || len(u.RongToken) == 0 {
		str, err := rcServer.UserGetToken(strconv.Itoa(u.Id), u.Nick, u.Avatar)
		if err == nil {
			var data struct {
				Code   int
				Token  string
				UserId string `json:"userId"`
			}
			json.Unmarshal(str, &data)
			u.RongUser = strconv.Itoa(u.Id)
			u.RongToken = data.Token
			o.Update(u)
		}
	}
	return true
}

func GetUser(uid int) (u *User, err error) {
	u = new(User)
	u.Id = uid
	err = orm.NewOrm().Read(u)
	return
}

func GetUserColumns(uid int) (list []*Column, likelist []*Column, err error) {
	u, err := GetUser(uid)
	if err != nil {
		return
	}
	_, err = orm.NewOrm().LoadRelated(u, "Columns")
	if err == nil {
		list = u.Columns
	}
	_, err = orm.NewOrm().LoadRelated(u, "LikeColumns")
	if err == nil {
		likelist = u.LikeColumns
	}
	return
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {

		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
