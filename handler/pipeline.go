package handler

import (
	"net/http"

	"gopkg.in/macaron.v1"
)

// Pipeline
func V1POSTPipelineHandler(ctx *macaron.Context) (int, []byte) {
	return http.StatusOK, nil
}
