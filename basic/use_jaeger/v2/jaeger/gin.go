package jaeger

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

var (
	SpanCTX = "span-ctx"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer := opentracing.GlobalTracer()
		var span opentracing.Span

		// 从header中获取jaeger信息
		wireContext, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		// 如果失败，构建新的span
		if err != nil {
			span = opentracing.StartSpan(c.FullPath())
		} else {
			span = opentracing.StartSpan(c.FullPath(), ext.RPCServerOption(wireContext))
		}

		defer span.Finish()

		c.Set(SpanCTX, opentracing.ContextWithSpan(c, span))
		c.Next()
	}
}

func StartSpan(tracer opentracing.Tracer, name string) opentracing.Span {
	//设置顶级span
	span := tracer.StartSpan(name)
	return span
}

func WithSpan(ctx context.Context, name string) (opentracing.Span, context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, name)
	return span, ctx
}

func GetCarrier(span opentracing.Span) (opentracing.HTTPHeadersCarrier, error) {
	carrier := opentracing.HTTPHeadersCarrier{}
	err := span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders, carrier)
	if err != nil {
		return nil, err
	}
	return carrier, nil
}

func GetParentSpan(spanName string, traceId string, header http.Header) (opentracing.Span, error) {
	carrier := opentracing.HTTPHeadersCarrier{}
	carrier.Set("uber-trace-id", traceId)

	tracer := opentracing.GlobalTracer()
	wireContext, err := tracer.Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(header),
	)

	parentSpan := opentracing.StartSpan(
		spanName,
		ext.RPCServerOption(wireContext))
	if err != nil {
		return nil, err
	}
	return parentSpan, err
}