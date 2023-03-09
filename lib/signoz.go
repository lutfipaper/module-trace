package lib

import (
	"context"
	"log"

	"github.com/lutfipaper/module-trace/interfaces"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc/credentials"
)

type SignozOpenTelemetry struct {
	option   interfaces.Option
	ctx      context.Context
	provider *sdktrace.TracerProvider
}

func NewSignozOpenTelemetry(option interfaces.Option) *SignozOpenTelemetry {
	return &SignozOpenTelemetry{option: option,
		ctx: context.Background()}
}

func (c *SignozOpenTelemetry) Setup() (err error) {
	secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if len(c.option.Config.Signoz.Insecure) > 0 {
		secureOption = otlptracegrpc.WithInsecure()
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(c.option.Config.Signoz.Endpoint),
		),
	)

	if err != nil {
		log.Fatal(err)
	}
	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", c.option.Name),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		return err
	}

	c.provider = sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resources),
	)
	otel.SetTracerProvider(c.provider)

	return err
}

func (c *SignozOpenTelemetry) Closing() (err error) {
	if c.provider != nil {
		err = c.provider.Shutdown(c.ctx)
	}
	return err
}
