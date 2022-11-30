package main

import (
	"shequn1/foundation/engine"
	"shequn1/foundation/server"
	"shequn1/internal/admin"
)

func main() {
	//go:generator
	server.Mode = "admin"
	//pages.New(&Table).List(ctx)
	server.Run(func(engine engine.Engine) { engine.Run(admin.GetRouter) })
}
