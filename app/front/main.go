package main

import (
	"shequn1/foundation/engine"
	"shequn1/foundation/server"
	"shequn1/internal/front"
)

func main() {
	server.Mode = "front"
	// pages.New(&Table).List(ctx)
	server.Run(func(engine engine.Engine) { engine.Run(front.GetRouter) })
}
