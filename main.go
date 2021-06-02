package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	_ "embed"

	"github.com/gin-gonic/gin"
)

//go:embed res/NotoSansCJKkr-Regular.otf
var fontfile []byte

//go:embed res/default.png
var defaultImg []byte

var fontSize = 80

func main() {
	dir, err := ioutil.TempDir("", "srvogimg")
	if err != nil {
		log.Fatal(err)
	}
	cachePath := os.Getenv("IMG_CACHE_PATH")
	if cachePath != "" {
		dir = cachePath
		os.Mkdir(cachePath, 0755)
	}

	r := gin.Default()
	r.GET("/render", func(c *gin.Context) {

		// Params
		text := c.DefaultQuery("text", "Hello, World!")
		imgurl := c.DefaultQuery("imgurl", "")
		startColor := c.DefaultQuery("startcolor", "E95420")
		endColor := c.DefaultQuery("endcolor", "772953")
		filepath := filepath.Join(dir, getHashedFileName(text+imgurl+startColor+endColor))

		_, err := ioutil.ReadFile(filepath)
		if err == nil {
			c.File(filepath)
		}

		path := drawOgImage(text, imgurl, startColor, endColor, filepath)
		c.File(path)
	})
	r.Run()
}
