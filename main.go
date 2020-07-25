package main

import (
	"github.com/gogf/gf/frame/g"
	_ "kong1/boot"
	_ "kong1/router"
)

func main() {
	g.Server().Run()
}
