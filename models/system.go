package models

import (
	"beego-demo/models/mymongo"
	"fmt"
	"gopkg.in/mgo.v2"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type System struct {
	ID         string `bson:"_id"`
	SystemCode string `bson:"systemcode"`
	SystemName string `bson:"systemname"`
}

type SystemForm struct {
	SystemCode string
	SystemName string
}

func NewSystem(s *SystemForm) (newp *System, err error) {

	system := System{
		SystemCode: s.SystemCode,
		SystemName: s.SystemName,
	}

	return &system, nil
}

func GetAllSystem() []System {
	mConn := mymongo.Conn()
	defer mConn.Close()

	var systemAll []System
	c := mConn.DB("").C("system")
	c.Find(nil).All(&systemAll)

	for i := 0; i < len(systemAll); i++ {
		fmt.Println("System ", systemAll[i].SystemCode, systemAll[i].SystemName)
		// }
	}
	return systemAll
}

func (s *System) Insert() (err error) {
	mConn := mymongo.Conn()
	defer mConn.Close()

	c := mConn.DB("").C("system")

	var u *System
	err = c.Find(bson.M{"systemcode": s.SystemCode}).One(u)
	if u !=nil {
		return errors.New("系统编码已经存在");
	}

	err = c.Insert(s)

	if err != nil {
		if mgo.IsDup(err) {
			return errors.New("系统编码已经存在");
		} else {
			return errors.New("数据库操作失败");
		}
	}
	return
}

func GetSystem(code string) (u *System, err error) {

	mConn := mymongo.Conn()
	defer mConn.Close()

	var system System
	c := mConn.DB("").C("system")
	err = c.Find(bson.M{"systemcode": code}).One(&system)
	if err == nil {
		u = &system
		return
	}

	return nil, errors.New("System not exists")
}
