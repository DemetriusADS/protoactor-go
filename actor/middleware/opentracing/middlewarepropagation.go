package opentracing

import (
	"github.com/DemetriusADS/protoactor-go/actor"
	"github.com/DemetriusADS/protoactor-go/actor/middleware/propagator"
)

// TracingMiddleware sets up spawn, sender, and receiver middlewares that propagate tracing spans.
func TracingMiddleware() actor.SpawnMiddleware {
	return propagator.New().
		WithItselfForwarded().
		WithSpawnMiddleware(SpawnMiddleware()).
		WithSenderMiddleware(SenderMiddleware()).
		WithReceiverMiddleware(ReceiverMiddleware()).
		SpawnMiddleware
}
