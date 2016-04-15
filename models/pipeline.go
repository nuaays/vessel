package models

type Pipeline struct {
	Id          int64  `json:"id"`
	WorkspaceId int64  `json:"workspaceId"`
	ProjectId   int64  `json:"projectId"`
	Name        string `json:"name" gorm:"type:varchar(255)"`
	SelfLink    string `json:"selfLink" gorm:"type:varchar(255)"`
	Created     int64  `json:"created"`
	Updated     int64  `json:"updated"`
	Labels      string `json:"labels"`
	Annotations string `json:"annotations"`
	Detail      string `json:"detail" gorm:"type:text"`
	Stages      []*Stage
	MetaData    PipelineMetaData
	StageSpecs  []StageSpec
}

type PipelineVersion struct {
	Id            int64    `json:"id"`
	WorkspaceId   int64    `json:"workspaceId"`
	ProjectId     int64    `json:"projectId"`
	PipelineId    int64    `json:"pipelineId"`
	Namespace     string   `json:"namespace"`
	SelfLink      string   `json:"selfLink" gorm:"type:varchar(255)"`
	Created       int64    `json:"created"`
	Updated       int64    `json:"updated"`
	Labels        string   `json:"labels"`
	Annotations   string   `json:"annotations"`
	Detail        string   `json:"detail" gorm:"type:text"`
	StageVersions []string `json:"stageVersions"`
	Log           string   `json:"log" gorm:"type:text"`
	Status        int64    `json:"state"`
	MetaData      PipelineMetaData
	StageSpecs    []StageSpec
}
