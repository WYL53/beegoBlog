package models

import (
//	"os"
//	"path"
	"strconv"
	"time"
	"github.com/go-xorm/xorm"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `xorm:"index"`
	Views           int64     `xorm:"index"`
	TopicTime       time.Time `xorm:"index"`
	TopicCount      int64
	TopicLassUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Category        string
	Content         string `xorm:"size(5000)"`
	Attachment      string
	Created         time.Time `xorm:"index"`
	Updated         time.Time `xorm:"index"`
	Views           int64     `xorm:"index"`
	Author          string
	ReplyTime       time.Time `xorm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func (this *Topic) IdString()string {
	return strconv.FormatInt(this.Id,10)
}


func registerDB() {
	engine.Sync2(new(Category),new(Topic))
	//orm.RegisterModel(new(Category), new(Topic))
	//orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	//orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddCategory(name string) error {
	//o := orm.NewOrm()
	cate := &Category{Title: name, Created: time.Now(), TopicTime: time.Now()}
	_,err := engine.Insert(cate)
	return err
/*
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
*/
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	//o := orm.NewOrm()
	cate := &Category{Id: cid}
	//_, err = o.Delete(cate)
	_,err = engine.Delete(cate)
	return err
}

func GetAllCategories() ([]*Category, error) {
	//o := orm.NewOrm()
	cates := make([]*Category, 0)
	//qs := o.QueryTable("category")
	//_, err := qs.All(&cates)
	err := engine.Find(&cates)
	return cates, err
}

func AddTopic(title, category, content string) error {
	//o := orm.NewOrm()
	topic := &Topic{
		Title:     title,
		Category:  category,
		Content:   content,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
	}
	//_, err := o.Insert(topic)
	_,err := engine.Insert(topic)
	return err
}

func GetAllTopics(isDesc bool) ([]*Topic, error) {
	//o := orm.NewOrm()
	topics := make([]*Topic, 0)
	//qs := o.QueryTable("topic")
	var sess *xorm.Session
	if isDesc {
		sess = engine.Desc("created")
	} else {
		sess = engine.Asc("created")
	}
	err := sess.Find(&topics)
	return topics, err
}

func GetTopic(tid string) (*Topic, error) {
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	//o := orm.NewOrm()
	topic := new(Topic)
	//qs := o.QueryTable("topic")
	//err = qs.Filter("id", id).One(topic)
	err = engine.Where("topic.id = ",id).Find(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = engine.Update(topic)
	return topic, err
}

func ModifyTopic(tid, title, category, content string) error {
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	//o := orm.NewOrm()
	topic := &Topic{
		Id: id,
	}
	ok,err := engine.Get(topic)
	//err = o.Read(topic)
	if ok && err == nil {
		topic.Title = title
		topic.Category = category
		topic.Content = content
		topic.Updated = time.Now()
		engine.Update(topic)
	}
	return err
}

func DeleteTopic(tid string) error {
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	//o := orm.NewOrm()
	topic := &Topic{Id: id}
	_, err = engine.Delete(topic)
	return err
}
