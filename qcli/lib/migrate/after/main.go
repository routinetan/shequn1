package main

import (
	"github.com/gogf/gf/container/glist"
	"github.com/gogf/gf/container/gset"
)

func SkipFile(src, dst *gset.Set) *glist.List {
	return glist.NewFrom(src.Diff(dst).Slice())
}
func SaveMigrateResult() []string {
	return []string{}
}

func main() {

}
