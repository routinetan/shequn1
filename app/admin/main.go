package main

import (
	"quanzi1/foundation/engine"
	"quanzi1/foundation/server"
	"quanzi1/internal/admin"
)

func main() {
	//go:generator
	server.Mode = "admin"
	// pages.New(&Table).List(ctx)
	server.Run(func(engine engine.Engine) { engine.Run(admin.GetRouter) })
}
