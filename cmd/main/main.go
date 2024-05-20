package main

import (
	_ "github.com/ZEQUANR/zhulong/driver"
	_ "github.com/ZEQUANR/zhulong/logger"
	"github.com/ZEQUANR/zhulong/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept", "X-Requested-With"}
	config.AllowOrigins = []string{"*"}
	config.ExposeHeaders = []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Cache-Control", "Content-Language", "Content-Type"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	routers.Init(r)

	r.Run()
}
