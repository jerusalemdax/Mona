package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Post struct {
	Id       int64
	Userid   int64  `orm:"index"`
	Author   string `orm:"size(15)"`
	Title    string `orm:"size(100)"`
	Color    string `orm:"size(7)"`
	Urlname  string `orm:"size(100);index"`
	Urltype  int8
	Content  string    `orm:"type(text)"`
	Tags     string    `orm:"size(100)"`
	Posttime time.Time `orm:"type(datetime);index"`
	Views    int64
	Status   int8
	Updated  time.Time `orm:"type(datetime)"`
	Istop    int8
	Cover    string `orm:"size(70)"`
}

func (m *Post) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
