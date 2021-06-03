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
	url := ginSwagger.URL("//" + hostname + "/api/swagger/doc.json")
	g.GET("/render", renderBasicImage)
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

// ShowAccount godoc
// @Summary OGP 미리보기 이미지 생성
// @Description OGP 미리보기 이미지 생성
// @ID render-ogp-image
// @Produce  image/png
// @Param text query string false "표시할 텍스트(중간에 공백이 있어야 자동 줄바꿈 처리됨)"
// @Param imgurl query string false "중앙 상단에 표시할 아이콘 이미지 파일의 웹 URL(없으면 표시 안하고, 이미지 조회중 오류 발생시 기본 이미지 표시)"
// @Param startcolor query string false "배경 그라데이션 색상값(좌측 상단 시작, 6자리 16진수 색상 코드 # 없이 입력)"
// @Param endcolor query string false "배경 그라데이션 색상값(우측 하단 끝, 6자리 16진수 색상 코드 # 없이 입력)"
// @Success 200
// @Router /api/render [get]
func renderBasicImage(c *gin.Context) {
	// Params
	text := c.DefaultQuery("text", "Hello, World!")
	imgurl := c.DefaultQuery("imgurl", "")
	startColor := c.DefaultQuery("startcolor", "030303")
	endColor := c.DefaultQuery("endcolor", "676767")
	filepath := filepath.Join(res.CachePath, generateHashFromString(text+imgurl+startColor+endColor))

	// Serve cached file if exists
	_, err := ioutil.ReadFile(filepath)
	if err == nil {
		c.File(filepath)
	}

	path := drawOgImage(text, imgurl, startColor, endColor, filepath)
	c.File(path)
}
