package draw

import (
	_ "embed"
	"image"
	"log"

	"github.com/sukso96100/srvogimg/res"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func drawBasicOgImage(
	text string,
	imgurls []string,
	logoimgurl string,
	bgimgurl string,
	bgStartColor string,
	bgEndColor string,
	isDarkTheme bool,
	filepath string) string {

	dc := gg.NewContext(ogImgWidth, ogImgHeight)
	userColors := res.GetThemeColors(isDarkTheme)

	if bgimgurl != "" {
		bgimg, err := LoadResizedBackgroundImage(bgimgurl)
		if err == nil {
			//Background image
			dc.DrawImage(bgimg, 0, 0)

			// Background image dimming
			dc.DrawRectangle(0, 0, ogImgWidthFloat, ogImgHeightFloat)
			if bgStartColor != "" {
				bgcolor, _ := ParseHexColor(bgStartColor, 200)
				dc.SetColor(bgcolor)
			} else {
				dc.SetColor(userColors.BackgroundColor)
			}
			dc.Fill()
		}
	} else {
		// Background gradient
		startColor := bgStartColor
		endColor := bgEndColor
		if startColor == "" {
			startColor = res.DefaultGradientStartColor
		}
		if endColor == "" {
			endColor = res.DefaultGradientEndColor
		}
		dc.DrawRectangle(0, 0, ogImgWidthFloat, ogImgHeightFloat)
		dc.SetFillStyle(NewGradientBackground(startColor, endColor, 255))
		dc.Fill()
	}

	resizedImgs := []image.Image{}
	widthTotal := 0
	padding := 20

	// Load Logo Images
	for _, imgurl := range imgurls {
		if imgurl != "" {
			m, _ := LoadResizedLogoImage(imgurl, 0, 200)
			resizedImgs = append(resizedImgs, m)
			widthTotal += m.Bounds().Size().X + padding
		}
	}

	// Draw Logo Images
	logoImgAnchor := (ogImgWidth / 2) - (widthTotal / 2)
	for _, img := range resizedImgs {
		logoImgAnchor += (img.Bounds().Size().X)/2 + (padding / 2)
		dc.DrawImageAnchored(img, logoImgAnchor, ogImgHeight*1/3, 0.5, 0.5)
		logoImgAnchor += (img.Bounds().Size().X)/2 + (padding / 2)
	}

	// Text
	f, err := opentype.Parse(res.Fontfile)
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}

	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    60,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	dc.SetFontFace(face)
	dc.SetColor(userColors.TextColor)
	if logoimgurl != "" {
		dc.DrawStringWrapped(text, ogImgWidthFloat/2, ogImgHeightFloat*3/5, 0.5, 0.5, 1000, 0.8, gg.AlignCenter)
	} else {
		dc.DrawStringWrapped(text, ogImgWidthFloat/2, ogImgHeightFloat*3/4, 0.5, 0.5, 1000, 0.8, gg.AlignCenter)
	}

	// Calculate Logo + Site name row width
	logoRowWidth := 0
	var logoimg image.Image
	var logoimgError error
	if logoimgurl != "" {
		logoimg, logoimgError = LoadResizedLogoImage(logoimgurl, 0, 80)
		if logoimgError == nil {
			logoRowWidth += logoimg.Bounds().Size().X
		}
	}

	// Draw Logo Images
	logoRowItemAnchor := (ogImgWidth / 2) - (logoRowWidth / 2)
	if logoimgurl != "" {
		logoRowItemAnchor += (logoimg.Bounds().Size().X) / 2
		dc.DrawImageAnchored(logoimg, logoRowItemAnchor, ogImgHeight*7/8, 0.5, 0.5)
		logoRowItemAnchor += (logoimg.Bounds().Size().X) / 2
	}

	imgfilePath := filepath + ".png"
	dc.SavePNG(imgfilePath)
	return imgfilePath
}
