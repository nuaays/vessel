package handler

import (
	"net/http"

	"gopkg.in/macaron.v1"
)

//UserBase
func V1POSTUserHandler(ctx *macaron.Context, stage StagePOSTJSON) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1PUTUserHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1GETUserHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

func V1DELETEUserHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}

//User auth
func V1UserSignInHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}
func V1UserSignOutHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, []byte("")
}
