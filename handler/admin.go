package handler

import (
	"net/http"

	"gopkg.in/macaron.v1"
)

//OrganizationBase
func V1POSTOrganizationHandler(ctx *macaron.Context, stage StagePOSTJSON) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1PUTOrganizationHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1GETOrganizationHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1DELETEOrganizationHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

//TeamBase
func V1POSTTeamHandler(ctx *macaron.Context, stage StagePOSTJSON) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1PUTTeamHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1GETTeamHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1DELETETeamHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}


//User to Organization
func V1POSTOrganizationUserHandler(ctx *macaron.Context, stage StagePOSTJSON) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1PUTOrganizationUserHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1GETOrganizationUserHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1DELETEOrganizationUserHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

//User to Team
func V1POSTTeamUserHandler(ctx *macaron.Context, stage StagePOSTJSON) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1PUTTeamUserHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1GETTeamUserHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1DELETETeamUserHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

//Team to Pipeline
func V1POSTPipelineTeamHandler(ctx *macaron.Context, stage StagePOSTJSON) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1PUTPipelineTeamHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1GETPipelineTeamHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1DELETEPipelineTeamHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}
