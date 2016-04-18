package models

// const(
// StageRun struct {
// 	Status
//
// }

)

type Stage struct {
	Id          int64    `json:"id"`
	WorkspaceId int64    `json:"workspaceId"`
	ProjectId   int64    `json:"projectId"`
	PipelineId  int64    `json:"pipelineId"`
	Created     int64    `json:"created"`
	Updated     int64    `json:"updated"`
	Name        string   `json:"name"`
	Detail      string   `json:"detail"`
	From        []string `josn:"from"`
	To          []string `json:"to"`
	MetaData    PipelineMetaData
	StageSpec   StageSpec
}

type StageVersion struct {
	Id                int64    `json:"id"`
	PipelineId        int64    `json:"pipelineId"`
	PipelineVersionId int64    `json:"pipelineVersionId"`
	StageId           int64    `json:"stageId"`
	Created           int64    `json:"created"`
	Updated           int64    `json:"updated"`
	Name              string   `json:"name"`
	Detail            string   `json:"detail"`
	State             int64    `json:"state"` // 0 not start    1 working    2 success    3 failed
	From              []string `josn:"from"`
	To                []string `json:"to"`
	MetaData          PipelineMetaData
	StageSpec         StageSpec
}

type StageVersionState struct {
	PipelineId        string `json:"pipelineId"`
	PipelineVersionId string `json:"pipelineVersionId"`
	StageId           string `json:"pipelineVersionId"`
	StageVersionId    string `json:"stageVersionId"`
	StageName         string `json:"stageName"`
	RunResult         string `json:"runResult"`
	Detail            string `json:"detail"`
}

func (stage *Stage) Create(pipelineId int64, name string) (error, string) {
	return nil, ""
}

func (stage *Stage) AddFrom(uuid string, from ...string) error {
	return nil
}

func (stage *Stage) AddEnd(uuid string, end ...string) error {
	return nil
}

func (stage *Stage) Run(uuid string) (error, string) {
	return nil, ""
}

func (stage *Stage) Copy(uuid string) (error, string) {
	return nil, ""
}
