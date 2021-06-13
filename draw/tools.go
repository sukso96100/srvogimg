package draw

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"image"
	"image/color"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/sukso96100/srvogimg/res"
)

const ogImgWidth int = 1200
const ogImgHeight int = 600
const ogImgWidthFloat float64 = 1200
const ogImgHeightFloat float64 = 600

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

func ParseHexColor(hex string, alpha uint8) (color.Color, error) {
	values, err := strconv.ParseUint(hex, 16, 32)

	if err != nil {
		return color.RGBA{0, 0, 0, 0}, err
	}
	red := uint8(values >> 16)
	blue := uint8((values >> 8) & 0xFF)
	green := uint8(values & 0xFF)

	rgb := color.RGBA{red, blue, green, alpha}

	return rgb, nil
}

func LoadResizedBackgroundImage(bgimgurl string) (image.Image, error) {
	resp, err := http.Get(bgimgurl)
	var iconimg []byte
	if err == nil {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			iconimg = body
			img, _, err := image.Decode(bytes.NewReader(iconimg))
			if err != nil {
				log.Fatalln(err)
			}
			dstImageFill := imaging.Fill(img, ogImgWidth, ogImgHeight, imaging.Center, imaging.Lanczos)
			return dstImageFill, nil
		}
		return nil, err
	}
	return nil, err
}

func LoadResizedLogoImage(imgurl string, width int, height int) (image.Image, error) {
	resp, err := http.Get(imgurl)
	var iconimg []byte
	if err != nil {
		// handle error
		iconimg = res.DefaultLogo
	} else {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			iconimg = body
		}
	}
	img, _, err := image.Decode(bytes.NewReader(iconimg))
	if err != nil {
		log.Fatalln(err)
	}
	dstImageFill := imaging.Resize(img, width, height, imaging.Lanczos)
	return dstImageFill, nil
}

func NewGradientBackground(start string, end string, alpha uint8) gg.Gradient {
	startRGB, _ := ParseHexColor(start, alpha)
	endRGB, _ := ParseHexColor(end, alpha)
	grad := gg.NewLinearGradient(0, 0, 1200, 600)
	grad.AddColorStop(0, startRGB)
	grad.AddColorStop(1, endRGB)
	return grad
}

func GenerateHashFromString(filename string) string {
	h := sha256.New()
	h.Write([]byte(filename))
	hbytes := h.Sum(nil)
	return fmt.Sprintf("%x", hbytes)
}

func GetThemeColors(isDarkTheme bool) ThemeColors {
	if isDarkTheme {
		return DarkColorTheme
	}
	return LightColorTheme
}
