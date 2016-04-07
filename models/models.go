package models

import (
	"os"
	"path"
	"strconv"
	"time"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLassUserId int64
}

const (
	_DB_NAME        = "data/beegoblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Category        string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{Title: name, Created: time.Now(), TopicTime: time.Now()}

	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return nil
	}
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func AddTopic(title, category, content string) error {
	o := orm.NewOrm()
	topic := &Topic{
		Title:     title,
		Category:  category,
		Content:   content,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
	}
	_, err := o.Insert(topic)
	return err
}

func GetAllTopics(isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	var err error
	if isDesc {
		_, err = qs.OrderBy().All(&topics)
	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

func GetTopic(tid string) (*Topic, error) {
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", id).One(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)
	return topic, err
}

func ModifyTopic(tid, title, category, content string) error {
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{
		Id: id,
	}
	err = o.Read(topic)
	if err == nil {
		topic.Title = title
		topic.Category = category
		topic.Content = content
		topic.Updated = time.Now()
		o.Update(topic)
	}
	return err
}

func DeleteTopic(tid string) error {
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id: id}
	_, err = o.Delete(topic)
	return err
}
