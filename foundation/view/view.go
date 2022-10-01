package view

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gview"
	"quanzi1/foundation/app"
)

var View *gview.View

type viewConfig struct {
	AutoEncode bool     `toml:"autoEncode"`
	Paths      []string `toml:"paths"`
	Delimiters []string `toml:"delimiters"`
}

var (
	viewconfig viewConfig
)

func Init() {
	err := app.Config().Bind("application", "view", &viewconfig)
	if err != nil {
		fmt.Println(err)
	}
	View = gview.New()
	View.SetConfigWithMap(g.Map{
		"Delimiters": viewconfig.Delimiters,
		"autoEncode": viewconfig.AutoEncode,
		"Paths":      viewconfig.Paths,
	})
}
