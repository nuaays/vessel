package models

import "github.com/jinzhu/gorm"

const (
	PipeLineType_K8S = 100
	PipeLineType_VM  = 200
	PipeLineType_PM  = 300

	PipeLineVersionType_K8S = 100
	PipeLineVersionType_VM  = 200
	PipeLineVersionType_PM  = 300

	StageType_K8S = 100
	StageType_VM  = 200
	StageType_PM  = 300

	PipelineVersionRunStatusSuccess = 200
	PipelineVersionRunStatusError   = 400

	PipelineVersionStageRunStatusSuccess = 200
	PipelineVersionStageRunStatusError   = 400
)

type Pipeline struct {
	gorm.Model
	Created      int64  `json:"created"`
	Updated      int64  `json:"updated"`
	Name         string `json:"name" gorm:"type:varchar(512)"`
	Actived      int64  `json:"actived"`
	PipelineType int64  `json:"pipelineType"`
	Description  string `json:"description"`
}

type PipelineVersion struct {
	gorm.Model
	Created             int64  `json:"created"`
	Updated             int64  `json:"updated"`
	Name                string `json:"name" gorm:"type:varchar(512)"`
	Actived             int64  `json:"actived"`
	PipelineVersionType int64  `json:"pipelineVersionType"`
	SpecDetail          string `json:"specDetail" gorm:"type:text"`
	Description         string `json:"description"`
}

type Stage struct {
	gorm.Model
	Created           int64    `json:"created"`
	Updated           int64    `json:"updated"`
	PipelineVersionId int64    `json:"pipelineVersionId"`
	Name              string   `json:"name" gorm:"type:varchar(512)"`
	From              []string `json:"from"`
	To                []string `json:"to"`
	StageType         int64    `json:"stageType"`
	SpecDetail        string   `json:"specDetail" gorm:"type:text"`
	Description       string   `json:"description"`
}

type PipelineVersionRun struct {
	gorm.Model
	Created           int64  `json:"created"`
	Updated           int64  `json:"updated"`
	PipelineVersionId int64  `json:"pipelineVersionId"`
	Status            int64  `json:"state"`
	Description       string `json:"description"`
}

type StageRun struct {
	gorm.Model
	Created              int64  `json:"created"`
	Updated              int64  `json:"updated"`
	PipelineVersionRunId int64  `json:"pipelineVersionRunId"`
	StageId              int64  `json:"stageId"`
	Status               int64  `json:"state"`
	RunResult            string `json:"runResult" gorm:"type:text"`
}
