package middleware

import (

	"gopkg.in/macaron.v1"
)

func InitLog(runmode, path string) {
	if runmode == "dev" {
	}

}

func logger(runmode string) macaron.Handler {
	return func(ctx *macaron.Context) {
		if runmode == "dev" {
		}
	}
}
