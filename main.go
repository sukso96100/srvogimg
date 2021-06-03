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
// @description A Service that renders Open Graph Protocol image to share on social media

func main() {

	res.InitCachePath()

	r := gin.Default()

	draw.SetupApis(r)
	r.Run()
}
