package models

import (
	"beego-demo/models/mymongo"
	"gopkg.in/mgo.v2/bson"
	"errors"
	"gopkg.in/mgo.v2"
	"time"
)

type RecordForm struct {
	//系统编码
	SystemCode string
	//项目名称
	ProjectName string
	//Review信息
	Message string
	//状态
	Status string
	//文件全名
	FileFullName string
	//行数
	RowNum int
	//列数
	ColumnNum int
}

type Record struct {
	RecordId int64 `bson:"recordid"`
	//系统编码
	SystemCode string `bson:"systemcode"`
	//项目名称
	ProjectName string `bson:"projectname"`
	//Review信息
	Message string `bson:"message"`
	//状态
	Status string `bson:"status"`
	//文件全名
	FileFullName string `bson:"filefullname"`
	//行数
	RowNum int `bson:"rownum"`
	//列数
	ColumnNum int `bson:"columnnum"`
}

type RecordListResult struct {
	CodeInfo
	Records []*Record
}

type RecordResult struct {
	CodeInfo
	RecordInfo *Record
}

func NewRecordResult(record *Record) (recordResult *RecordResult) {

	recordResult = &RecordResult{
		RecordInfo: record,
		CodeInfo: CodeInfo{
			Code: 1,
			Info: "success",
		},
	}
	return recordResult
}

func NewRecord(form *RecordForm) *Record {

	record := Record{
		RecordId:     time.Now().Unix(),
		SystemCode:   form.SystemCode,
		ProjectName:  form.ProjectName,
		Message:      form.Message,
		Status:       form.Status,
		FileFullName: form.FileFullName,
		RowNum:       form.RowNum,
		ColumnNum:    form.ColumnNum,
	}
	return &record
}

func (r *Record) AddRecord() (err error) {

	mConn := mymongo.Conn()
	defer mConn.Close()

	c := mConn.DB("").C("record")
	err = c.Insert(r)

	if err != nil {
		if mgo.IsDup(err) {
			return errors.New("系统编码已经存在");
		} else {
			return errors.New("数据库操作失败");
		}
	}
	return
}

func GetAllRecordByCode(SystemCode string) ([]*Record, error) {
	mConn := mymongo.Conn()
	defer mConn.Close()
	c := mConn.DB("").C("record")

	var recordAll []*Record
	err := c.Find(bson.M{"systemcode": SystemCode}).All(&recordAll)

	if err != nil {
		return nil, errors.New("查询记录失败");
	}

	return recordAll, nil

}

func GetRecordByID(RecordId int64) (r *Record, err error) {

	mConn := mymongo.Conn()
	defer mConn.Close()
	c := mConn.DB("").C("record")

	var record Record
	err = c.Find(bson.M{"recordid": RecordId}).One(&record)

	if err != nil {
		return nil, errors.New("查询记录失败");
	}

	r = &record
	return
}
