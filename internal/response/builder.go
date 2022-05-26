package response

import (
	"io"
	"net/http"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Headerer interface {
	Header() http.Header
}

type AppCoder interface {
	AppCode() int
}

// API 基本响应
type Response string

func Build(code int, message string, data interface{}) Response {
	return Response(`{}`).
		WithCode(code, message).
		WithData(data)
}

func FromError(err error) Response {
	// 未知错误
	code := 500

	if appCode, ok := err.(AppCoder); ok {
		code = appCode.AppCode()
	}

	return Build(code, err.Error(), nil)
}

func Make() Response {
	return Response(`{"code": 0, "message": ""}`)
}

func (r Response) MustSet(path string, value interface{}) Response {
	rv, _ := sjson.Set(string(r), path, value)
	return Response(rv)
}

func (r Response) MustSetRaw(path, value string) Response {
	rv, _ := sjson.SetRaw(string(r), path, value)
	return Response(rv)
}

func (r Response) MustSetData(path string, value interface{}) Response {
	return r.MustSet("data."+path, value)
}

func (r Response) WithCode(code int, message string) Response {
	return r.MustSet("code", code).MustSet("message", message)
}

func (r Response) WithData(data interface{}) Response {
	if data == nil {
		// nil 的时候不设置 data 字段
		return r
	}

	return r.MustSet("data", data)
}

func (r Response) WithDataRaw(dataStr string) Response {
	return r.MustSetRaw("data", dataStr)
}

func (r Response) String() string {
	return string(r)
}

func (r Response) MarshalJSON() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r Response) WriteTo(w io.Writer) (int64, error) {
	b, err := r.MarshalJSON()
	if err != nil {
		return 0, err
	}

	if wh, ok := w.(Headerer); ok {
		// NOTE: 对 HTTP 响应尝试设置 content type
		wh.Header().Set("Content-Type", "application/json; charset=utf-8")
	}

	wc, err := w.Write(b)
	return int64(wc), err
}

func (r Response) Get(path string) gjson.Result {
	return gjson.Get(string(r), path)
}
