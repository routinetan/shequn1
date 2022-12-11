package main

import (
	"shequn1/internal/foundation/engine"
	"shequn1/internal/foundation/server"
	"shequn1/internal/front"
)

func main() {
	server.Mode = "front"
	// pages.New(&Table).List(ctx)
	server.Run(func(engine engine.Engine) { engine.Run(front.GetRouter) })
}
