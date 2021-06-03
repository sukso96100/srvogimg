package main

import (
	_ "embed"

	_ "embed"

	"github.com/gin-gonic/gin"
	"github.com/sukso96100/srvogimg/draw"
	"github.com/sukso96100/srvogimg/res"
)

// @title srvogimg
// @version 1.0
// @description Open Graph Protocol Card image generater service

func main() {

	res.InitCachePath()

	r := gin.Default()

	draw.SetupApis(r)
	r.Run()
}
