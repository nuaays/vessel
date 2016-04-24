package handler

// BaseResp base responce for all responce
type BaseResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

const (
	// CodeSuccess :
	CodeSuccess = 200
	// CodeError :
	CodeError = 500

	// MessageSuccess :
	MessageSuccess = "success"
	// MessageError :
	MessageError = "error"
)
