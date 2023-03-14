package template

var InternalMysql = `package internal

import "github.com/go-xorm/xorm"

func MysqlConn() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:root@/example?charset=utf8")
	if err != nil {
		panic(err)
	}
	return engine
}`
