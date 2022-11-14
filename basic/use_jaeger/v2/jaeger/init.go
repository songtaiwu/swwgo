package jaeger

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"io/ioutil"
	"log"
)


var JaegerUrl = "192.168.71.131"
var jaegerPort = 6831

func StartJaeger(serviceName string) {
	tracer, _ := initJaeger(serviceName)
	opentracing.SetGlobalTracer(tracer)
}

func initJaeger(serviceName string) (opentracing.Tracer, io.Closer) {
	if JaegerUrl == "" {
		return opentracing.NoopTracer{}, ioutil.NopCloser(nil)
	}

	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type: "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort: fmt.Sprintf("%s:%d", JaegerUrl, jaegerPort),
			LogSpans: true,
		},
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.Fatalf("Failed to init Jaeger client: %s", err)
	}
	return tracer, closer
}