package handler

import (
	"net/http"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/macaron.v1"

	"github.com/containerops/vessel/models"
	"github.com/containerops/vessel/module/pipeline"
)

func V1POSTPipelineHandler(ctx *macaron.Context, reqData models.PipelineSpecTemplate) (int, []byte) {
	/*
		etcd path /vessel/ws-xxx/pj-xxx/pl-xxx/plv-xxx/stage-xxx/

		plv-xxx  -> k8s namespace

		demo:
		/containerops/vessel/ws-xxx/pj-xxx/pl-xxx1/stage/stage1/...
		/containerops/vessel/ws-xxx/pj-xxx/pl-xxx1/stage/stage2...
		/containerops/vessel/ws-xxx/pj-xxx/pl-xxx2/stage/stage1/...
		/containerops/vessel/ws-xxx/pj-xxx/pl-xxx2/stage/stage2...
		/containerops/vessel/ws-xxx/pj-xxx/pl-xxx1/plv-xxx/stagev-xxx/name
		/containerops/vessel/ws-xxx/pj-xxx/pl-xxx1/plv-xxx/stagev-xxx/dependence/Dependence1ServicesName <--need watch
		/containerops/vessel/ws-xxx/pj-xxx/pl-xxx1/plv-xxx/stagev-xxx/dependence/Dependence2ServicesName <--need watch
		/containerops/vessel/ws-xxx/pj-xxx/pl-xxx1/plv-xxx/stagev-xxx/check/check_status_url
		/containerops/vessel/ws-xxx/pj-xxx/pl-xxx1/plv-xxx/stagev-xxx/check/check_status_interval
		/containerops/vessel/ws-xxx/pj-xxx/pl-xxx1/plv-xxx/stagev-xxx/check/check_status_count
	*/

	pl, err := createPipelineAndStage(reqData)
	if err != nil {
		return http.StatusOK, []byte(err.Error())
	}

	plv, err := pipeline.RunPipeline(pl)
	if err != nil {
		return http.StatusOK, []byte(err.Error())
	}

	result := pipeline.BootPipelineVersion(plv)

	return http.StatusOK, []byte(result)
}

func createPipelineAndStage(pst models.PipelineSpecTemplate) (*models.Pipeline, error) {
	var plInfo *models.Pipeline = &models.Pipeline{}
	plInfo.Id = time.Now().UnixNano()
	//ignore workspace & project
	plInfo.WorkspaceId = 10000
	plInfo.ProjectId = 20000
	plInfo.Name = pst.MetaData.Name
	plInfo.Created = time.Now().Unix()
	plInfo.Updated = time.Now().Unix()
	// plInfo.Labels = pst.MetaData.Labels
	// plInfo.Annotations = pst.MetaData.Annotations
	//ignore plJson Detail
	plInfo.Detail = ""
	plInfo.MetaData = pst.MetaData
	plInfo.StageSpecs = pst.Spec

	for _, value := range pst.Spec {
		var sInfo *models.Stage = &models.Stage{}
		sInfo.Id = time.Now().Unix()
		sInfo.WorkspaceId = plInfo.WorkspaceId
		sInfo.ProjectId = plInfo.ProjectId
		sInfo.PipelineId = plInfo.Id
		sInfo.Created = time.Now().Unix()
		sInfo.Updated = time.Now().Unix()
		sInfo.Name = value.Name
		//ignore Stage Detail
		//StatusCheckUrl to Detail
		//StatusCheckInterval to Detail
		//StatusCheckCount to Detail
		sInfo.Detail = ""
		sInfo.From = strings.Split(value.Dependence, ",")
		if len(sInfo.From) == 1 && sInfo.From[0] == "" {
			sInfo.From = make([]string, 0)
		}
		sInfo.MetaData = plInfo.MetaData
		sInfo.StageSpec = value

		plInfo.Stages = append(plInfo.Stages, sInfo)
	}

	log.Error(plInfo)
	for _, value := range plInfo.Stages {
		log.Error(value)
	}

	return plInfo, nil
}

func V1PUTPipelineHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1GETPipelineHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1DELETEPipelineHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1RunPipelineHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1StatusGETHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}
