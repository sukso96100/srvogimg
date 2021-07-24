package draw

import (
	"io/ioutil"
	"path/filepath"
	"strconv"

	"github.com/sukso96100/srvogimg/res"

	_ "github.com/sukso96100/srvogimg/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupApis(g *gin.Engine) {
	url := ginSwagger.URL("doc.json")
	g.GET("/basic", renderBasicImage)
	g.GET("/article", renderArticleImage)
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

// ShowAccount godoc
// @Summary Render a Basic OGP image
// @Description Render a OGP image with logo and text
// @ID basic-ogp-image
// @Produce  image/png
// @Param text query string false "Text to display(Needs space for automatic linebreak)"
// @Param imgurl query string false "Web URL of the logo image to display(Show default image if error occured when loading)"
// @Param imgurl2 query string false "Web URL of the logo image to display(Show default image if error occured when loading)"
// @Param imgurl3 query string false "Web URL of the logo image to display(Show default image if error occured when loading)"
// @Param logoimgurl query string false "Web URL of the logo image"
// @Param bgimgurl query string false "Web URL of the background image"
// @Param bgstartcolor query string false "Background gradient start (top left) color(Color code in HEX without #)"
// @Param bgendcolor query string false "Background gradient end (bottom right) color(Color code in HEX without #)"
// @Param isdark query string false "Choose color scheme (true: dark, false: light)"
// @Success 200
// @Router /basic [get]
func renderBasicImage(c *gin.Context) {
	// Params
	text := c.DefaultQuery("text", "Hello, World!")
	imgurl := c.DefaultQuery("imgurl", "")
	imgurl2 := c.DefaultQuery("imgurl2", "")
	imgurl3 := c.DefaultQuery("imgurl3", "")
	logoimgurl := c.DefaultQuery("logoimgurl", "")
	bgimgurl := c.DefaultQuery("bgimgurl", "")
	startColor := c.DefaultQuery("bgstartcolor", "")
	endColor := c.DefaultQuery("bgendcolor", "")
	isDark := c.DefaultQuery("isdark", "true") == "true"

	filepath := filepath.Join(res.CachePath,
		GenerateHashFromString(text+imgurl+imgurl2+imgurl3+logoimgurl+startColor+endColor))

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

	path := drawBasicOgImage(text, imgurls, logoimgurl, bgimgurl, startColor, endColor, isDark, filepath)
	c.File(path)
}

// ShowAccount godoc
// @Summary Render a Article OGP image
// @Description Render a OGP image with Title, Author, website logo and website name.
// @ID article-ogp-image
// @Produce  image/png
// @Param title query string false "Title of article to display(Needs space for automatic linebreak)"
// @Param authors query string false "authors of article to display(Needs space for automatic linebreak)"
// @Param sitename query string false "Website name to display(Needs space for automatic linebreak)"
// @Param bgimgurl query string false "Web URL of the background image"
// @Param logoimgurl query string false "Web URL of the logo image"
// @Param bgstartcolor query string false "Background gradient start (top left) color(Color code in HEX without #)"
// @Param bgendcolor query string false "Background gradient end (bottom right) color(Color code in HEX without #)"
// @Param isdark query string false "Choose color scheme (true: dark, false: light)"
// @Success 200
// @Router /article [get]
func renderArticleImage(c *gin.Context) {
	title := c.DefaultQuery("title", "Hello, World!")
	authors := c.DefaultQuery("authors", "Author")
	sitename := c.DefaultQuery("sitename", "My Website")
	bgimgurl := c.DefaultQuery("bgimgurl", "")
	logoimgurl := c.DefaultQuery("logoimgurl", "")
	bgStartColor := c.DefaultQuery("bgstartcolor", "")
	bgEndColor := c.DefaultQuery("bgendcolor", "")
	isDark := c.DefaultQuery("isdark", "true") == "true"

	filepath := filepath.Join(res.CachePath,
		GenerateHashFromString(title+authors+sitename+bgimgurl+logoimgurl+bgStartColor+bgEndColor+strconv.FormatBool(isDark)))

	path := drawArticleOgImage(title, authors, sitename, bgimgurl, logoimgurl, bgStartColor, bgEndColor, isDark, filepath)
	c.File(path)
}
