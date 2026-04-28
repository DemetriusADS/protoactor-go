package opentelemetry

import (
	"github.com/DemetriusADS/protoactor-go/actor"
	"github.com/DemetriusADS/protoactor-go/actor/middleware/propagator"
)

func TracingMiddleware() actor.SpawnMiddleware {
	return propagator.New().
		WithItselfForwarded().
		WithSpawnMiddleware(SpawnMiddleware()).
		WithSenderMiddleware(SenderMiddleware()).
		WithReceiverMiddleware(ReceiverMiddleware()).
		SpawnMiddleware
}
