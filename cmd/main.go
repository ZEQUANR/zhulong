package main

import (
	_ "github.com/ZEQUANR/zhulong/driver"
	"github.com/ZEQUANR/zhulong/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.Init(r)

	r.Run()
}
