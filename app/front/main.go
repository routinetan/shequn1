package main

import (
	"fmt"
	"github.com/gogf/gf/os/gspath"
	"shequn1/foundation/engine"
	"shequn1/foundation/server"
	"shequn1/internal/front"
)

func main() {
	server.Mode = "front"
	path, _ := gspath.Search("/www/shequn1/web/", "/front/")
	fmt.Println(path)
	// pages.New(&Table).List(ctx)
	server.Run(func(engine engine.Engine) { engine.Run(front.GetRouter) })
}
