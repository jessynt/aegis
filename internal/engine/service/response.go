package service

import (
	"aegis/internal/model"
)

// TODO: generate with protoc

func (r *ReloadResponse) SetError(err error) {
	switch err := err.(type) {
	case nil:
		return
	case model.AppCoder:
		r.Code = int64(err.AppCode())
		r.ErrorReason = err.AppMessage()
		return
	default:
		r.Code = 500
		r.ErrorReason = err.Error()
	}
}

func (r *ReloadResponse) GetError() error {
	if r.Code == 0 {
		return nil
	}
	return model.MakeAppError(int(r.Code), r.ErrorReason)
}
