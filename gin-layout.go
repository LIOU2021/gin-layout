package main

import (
	"github.com/LIOU2021/gin-layout/config"
	"github.com/LIOU2021/gin-layout/core"
)

func init() {
	config.Kernel()
}

func main() {
	core.Run()
}
