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
