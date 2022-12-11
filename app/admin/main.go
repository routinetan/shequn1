package main

import (
	"shequn1/internal/admin"
	"shequn1/internal/foundation/engine"
	"shequn1/internal/foundation/server"
)

func main() {
	//go:generator
	server.Mode = "admin"
	//pages.New(&Table).List(ctx)
	server.Run(func(engine engine.Engine) { engine.Run(admin.GetRouter) })
}
