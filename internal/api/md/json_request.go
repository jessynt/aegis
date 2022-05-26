package md

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"

	"aegis/internal/response"
)

var (
	emptyJSONPayload = gjson.Parse("{}")
)

const (
	CONTEXT_KEY_JSON_PAYLOAD = "json_payload"
)

// 解析请求的 JSON 内容
//
// - forceParse: 是否强制解析，如果为假，将不会进行解析
// - required: JSON 内容是否必须项，如果为真，且 JSON 内容为空，直接返回 `BadRequest`
func PrepareJSONPayload(forceParse, required bool) Middleware {
	return func(next http.Handler) http.Handler {
		emptyPayload := func(w http.ResponseWriter, r *http.Request) {
			nextCtx := context.WithValue(r.Context(), CONTEXT_KEY_JSON_PAYLOAD, emptyJSONPayload)
			next.ServeHTTP(w, r.WithContext(nextCtx))
		}

		badRequest := func(w http.ResponseWriter) {
			_, _ = response.ResponseBadRequest.WriteTo(w)
		}

		f := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			if ctx.Value(CONTEXT_KEY_JSON_PAYLOAD) != nil {
				// 之前已经处理过了
				next.ServeHTTP(w, r)
				return
			}

			if r.ContentLength <= 0 {
				// 没有 body 可以解析
				if required {
					badRequest(w)
					return
				}

				next.ServeHTTP(w, r)
				return
			}

			if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
				// 不是标准的 JSON 请求
				if !forceParse {
					emptyPayload(w, r)
					return
				}
			}

			raw, err := ioutil.ReadAll(r.Body)
			if err != nil {
				// 解析失败
				if required {
					badRequest(w)
					return
				}
				emptyPayload(w, r)
				return
			}

			ctx = context.WithValue(ctx, CONTEXT_KEY_JSON_PAYLOAD, gjson.ParseBytes(raw))
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(f)
	}
}

var RequireJSONPayload = PrepareJSONPayload(true, true)

// 从 context 中获取 JSON 请求内容
func PopulateJSONPayload(ctx context.Context) gjson.Result {
	ivalue := ctx.Value(CONTEXT_KEY_JSON_PAYLOAD)
	if ivalue == nil {
		return emptyJSONPayload
	}

	if value, ok := ivalue.(gjson.Result); ok {
		return value
	} else {
		return emptyJSONPayload
	}
}
