package res

import (
	_ "embed"
	"image/color"
	"io/ioutil"
	"log"
	"os"
)

//go:embed font.otf
var Fontfile []byte

//go:embed default.png
var DefaultLogo []byte

var CachePath string

type ThemeColors struct {
	TextColor       color.Color
	BackgroundColor color.Color
}

var LightColorTheme ThemeColors = ThemeColors{
	color.Black,
	color.RGBA{174, 167, 159, 220},
}

var DarkColorTheme ThemeColors = ThemeColors{
	color.White,
	color.RGBA{0, 0, 0, 200},
}

func GetThemeColors(isDarkTheme bool) ThemeColors {
	if isDarkTheme {
		return DarkColorTheme
	}
	return LightColorTheme
}

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

var DefaultGradientStartColor string = "EB6536"
var DefaultGradientEndColor string = "772953"

func InitDefaultGradientColors() {
	if os.Getenv("DEFAULT_START_COLOR") != "" {
		DefaultGradientStartColor = os.Getenv("DEFAULT_START_COLOR")
	}
	if os.Getenv("DEFAULT_END_COLOR") != "" {
		DefaultGradientEndColor = os.Getenv("DEFAULT_END_COLOR")
	}
}
