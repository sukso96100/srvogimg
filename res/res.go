package res

import (
	_ "embed"
	"io/ioutil"
	"log"
	"os"
)

//go:embed font.otf
var Fontfile []byte

//go:embed default.png
var DefaultLogo []byte

var CachePath string

func InitCachePath() {
	CachePath = os.Getenv("IMG_CACHE_PATH")
	if CachePath != "" {
		os.Mkdir(CachePath, 0755)
	} else {
		dir, err := ioutil.TempDir("", "srvogimg")
		if err != nil {
			log.Fatal(err)
		}
		CachePath = dir
	}
}

var DefaultGradientStartColor string = "E95420"
var DefaultGradientEndColor string = "772953"

func InitDefaultGradientColors() {
	if os.Getenv("DEFAULT_START_COLOR") != "" {
		DefaultGradientStartColor = os.Getenv("DEFAULT_START_COLOR")
	}
	if os.Getenv("DEFAULT_END_COLOR") != "" {
		DefaultGradientEndColor = os.Getenv("DEFAULT_END_COLOR")
	}
}
