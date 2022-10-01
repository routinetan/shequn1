package main

import (
	"quanzi1/foundation/engine"
	"quanzi1/foundation/server"
	"quanzi1/internal/front"
)

func main() {
	server.Mode = "front"

	// pages.New(&Table).List(ctx)
	server.Run(func(engine engine.Engine) { engine.Run(front.GetRouter) })
}
