package trace

import (
	"github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
)

func ExtractTraceId(span opentracing.Span) string {
	jaegerSpan, ok := span.(*jaeger.Span)
	if !ok {
		return "N/A"
	}

	jaegerSpanCtx, ok := jaegerSpan.Context().(jaeger.SpanContext)
	if !ok {
		return "N/A"
	}

	return jaegerSpanCtx.TraceID().String()
}
