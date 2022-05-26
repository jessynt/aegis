package response

import (
	"aegis/internal/model"
)

var (
	ErrBadRequest = model.MakeAppError(400, "bad request")
	ErrNotFound   = model.MakeAppError(404, "not found")
)

var (
	ResponseNoContent  = Build(0, "", nil)
	ResponseBadRequest = FromError(ErrBadRequest)
	ResponseNotFound   = FromError(ErrNotFound)
)
