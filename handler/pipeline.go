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

// Pipeline
func V1POSTPipelineHandler(ctx *macaron.Context, reqData models.PipelineSpecTemplate) (int, []byte) {
	return http.StatusOK, nil
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

// PipelineVersion
func V1POSTPipelineVersionHandler(ctx *macaron.Context, reqData models.PipelineSpecTemplate) (int, []byte) {
	return http.StatusOK, nil
}

func V1GETPipelineVersionHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1DELETEPipelineVersionHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1RunPipelineVersionHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}
