package draw

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sukso96100/srvogimg/res"

	_ "github.com/sukso96100/srvogimg/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupApis(g *gin.Engine) {
	hostname := os.Getenv("APP_HOST")
	if hostname == "" {
		hostname = "localhost:8080"
	}
	url := ginSwagger.URL("//" + hostname + "/swagger/doc.json")
	g.GET("/render", renderBasicImage)
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

// ShowAccount godoc
// @Summary Render a OGP image
// @Description Render a OGP image with logo and text
// @ID render-ogp-image
// @Produce  image/png
// @Param text query string false "Text to display(Needs space for automatic linebreak)"
// @Param imgurl query string false "Web URL of the logo image to display(Show default image if error occured when loading)"
// @Param imgurl2 query string false "Web URL of the logo image to display(Show default image if error occured when loading)"
// @Param imgurl3 query string false "Web URL of the logo image to display(Show default image if error occured when loading)"
// @Param startcolor query string false "Background gradient start (top left) color(Color code in HEX without #)"
// @Param endcolor query string false "Background gradient end (bottom right) color(Color code in HEX without #)"
// @Success 200
// @Router /render [get]
func renderBasicImage(c *gin.Context) {
	// Params
	text := c.DefaultQuery("text", "Hello, World!")
	imgurl := c.DefaultQuery("imgurl", "")
	imgurl2 := c.DefaultQuery("imgurl2", "")
	imgurl3 := c.DefaultQuery("imgurl3", "")
	startColor := c.DefaultQuery("startcolor", res.DefaultGradientStartColor)
	endColor := c.DefaultQuery("endcolor", res.DefaultGradientEndColor)
	filepath := filepath.Join(res.CachePath,
		generateHashFromString(text+imgurl+imgurl2+imgurl3+startColor+endColor))

	// Serve cached file if exists
	_, err := ioutil.ReadFile(filepath)
	if err == nil {
		c.File(filepath)
	}

	imgurls := []string{}
	if imgurl != "" {
		imgurls = append(imgurls, imgurl)
	}
	if imgurl2 != "" {
		imgurls = append(imgurls, imgurl2)
	}
	if imgurl3 != "" {
		imgurls = append(imgurls, imgurl2)
	}

	path := drawBasicOgImage(text, imgurls, startColor, endColor, filepath)
	c.File(path)
}
