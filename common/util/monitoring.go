package util

import (
	"context"
	"net/http"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type SpanOption struct {
	Context     context.Context
	Entity      string
	Name        string
	HttpRequest *http.Request
}

func StartSpanFromContextRest(option SpanOption) (ddtrace.Span, context.Context) {
	return tracer.StartSpanFromContext(option.Context, "rest."+option.Entity+"."+option.Name,
		tracer.SpanType(ext.SpanTypeWeb),
		tracer.ResourceName(option.HttpRequest.Method+" "+option.HttpRequest.URL.Path),
		tracer.Tag(ext.HTTPMethod, option.HttpRequest.Method),
		tracer.Tag(ext.HTTPURL, option.HttpRequest.URL.Path),
	)
}

func StartSpanFromContextUsecase(ctx context.Context, entity, name string) (ddtrace.Span, context.Context) {
	return tracer.StartSpanFromContext(ctx, "usecase."+entity+"."+name)
}

func StartSpanFromContextConsumer(option SpanOption) (ddtrace.Span, context.Context) {
	return tracer.StartSpanFromContext(option.Context, "consumer."+option.Entity+"."+option.Name,
		tracer.SpanType(ext.SpanTypeMessageConsumer))
}
