package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/astaxie/beego"
	"fmt"
)

var engine *xorm.Engine

func init(){
	dbUser := beego.AppConfig.String("dbUser")
	dbPasswd := beego.AppConfig.String("dbPasswd")
	dbHost := beego.AppConfig.String("dbHost")
	dbName := beego.AppConfig.String("dbName")
	databaseUrl := fmt.Sprintf("%s:%s@%s/%s?charset=latin1",dbUser,dbPasswd,dbHost,dbName)
	fmt.Println("init database:"+databaseUrl)
	var err error
	engine, err = xorm.NewEngine("mysql",databaseUrl)
	if err != nil{
		panic("database connect faild.")
	}
	engine.SetMaxIdleConns(5) //连接池的空闲数大小
	engine.SetMaxOpenConns(10) //最大打开连接数
	fmt.Printf("%s connect success.",databaseUrl)
	err = engine.Ping()
	if err != nil{
		panic("database ping faild.")
	}
	registerDB()
}