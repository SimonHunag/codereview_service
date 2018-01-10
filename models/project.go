package models

import (
	"beego-demo/models/mymongo"
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
)

type Project struct {
	ProjectId   int64 `bson:"projectid"`
	ProjectName string `bson:"projectname"`
}

func GetAllProject() []Project {
	mConn := mymongo.Conn()
	defer mConn.Close()

	var personAll []Project
	c := mConn.DB("").C("project")
	c.Find(nil).All(&personAll)

	for i := 0; i < len(personAll); i++ {
		fmt.Println("Person ", personAll[i].ProjectId, personAll[i].ProjectName)
		// }
	}
	return personAll
}

func (p *Project) Insert() (code int64, err error) {
	mConn := mymongo.Conn()
	defer mConn.Close()

	c := mConn.DB("").C("project")
	err = c.Insert(p)

	if err != nil {
		if mgo.IsDup(err) {
			code = ErrDupRows
		} else {
			code = ErrDatabase
		}
	} else {
		code = p.ProjectId
	}
	return
}

func NewProject(p *Project) (newp *Project, err error) {

	project := Project{
		ProjectId:   time.Now().Unix(),
		ProjectName: p.ProjectName,
	}

	return &project, nil
}
