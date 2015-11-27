package models

import (
	"time"
)

const (
	PIPESUCCESS = iota
	PIPEERROR
	PIPERUNNING
	PIPEPENDDING
	PIPEEXCEPT
)

type Pipeline struct {
	Id          int64     `json:"id"`                                        //
	ProjectId   int64     `json:"projectId"`                                 //
	UUID        string    `json:"uuid" orm:"varchar(255)"`                   //
	Name        string    `json:"name" orm:"varchar(255)"`                   //
	Description string    `json:"description" orm:"null;type(text)"`         //
	Actived     bool      `json:"actived" orm:"null;default(0)"`             //
	RootId      int64     `json:"rootId" orm:"default(0)"`                   //
	ParentId    int64     `json:"parentId" orm:"default(0)"`                 //
	Version     bool      `json:"version" orm:"default(0)"`                  //
	Created     time.Time `json:"created" orm:"auto_now_add;type(datetime)"` //
	Updated     time.Time `json:"updated" orm:"auto_now;type(datetime)"`     //
	Memo        string    `json:"memo" orm:"null;type(text)"`                //
}

type Status struct {
	Id         int64     `json:"id"`                                        //
	PipelineId int64     `json:"pipelineId"`                                //
	UUID       string    `json:"uuid" orm:"varchar(255)"`                   //
	Resource   string    `json:"resrouce" orm:"null;type(text)"`            //
	ActionId   string    `json:"actionId" orm:"unique;varchar(255)"`        // Point.UUID; Stage.UUID
	Started    time.Time `json:"started" orm:"type(datetime)"`              //
	Ended      time.Time `json:"ended" orm:"type(datetime)"`                //
	Log        string    `json:"log" orm:"type(text)"`                      //
	Result     int64     `json:"result" orm:"null"`                         // Success: 0; Failure: 1
	Actived    bool      `json:"actived" orm:"null;default(0)"`             //
	Created    time.Time `json:"created" orm:"auto_now_add;type(datetime)"` //
	Updated    time.Time `json:"updated" orm:"auto_now;type(datetime)"`     //
	Memo       string    `json:"memo" orm:"null;type(text)"`                //
}

//Create Pipeline
func (pipe *Pipeline) Create(projectId int64, name string) (error, int64) {
	return nil, 0
}

//Create status records with same uuid
func (pipe *Pipeline) Run(pipelineId int64) (error, string) {
	return nil, ""
}

//List all run history with pipelineId
func (pipe *Pipeline) History(pipelineId int64) (error, []string) {
	return nil, []string{}
}

//Return all point and stage status with status uuid
func (pipe *Pipeline) Status(uuid string) (error, []int64) {
	return nil, []int64{}
}

//Stop run
func (pipe *Pipeline) Stop(uuid string) error {
	return nil
}

//Clean run resources
func (pipe *Pipeline) Clean(uuid string) error {
	return nil
}

func (pipe *Pipeline) Copy(uuid string) (error, string) {
	return nil, ""
}
