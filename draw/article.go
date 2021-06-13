package draw

import (
	_ "embed"
	"image"
	"log"

	"github.com/fogleman/gg"
	"github.com/sukso96100/srvogimg/res"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func drawArticleOgImage(
	title string,
	authors string,
	sitename string,
	bgimgurl string,
	logoimgurl string,
	bgStartColor string,
	bgEndColor string,
	isDarkTheme bool,
	filepath string) string {
	dc := gg.NewContext(ogImgWidth, ogImgHeight)
	userColors := GetThemeColors(isDarkTheme)

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

	// Title
	f, err := opentype.Parse(res.Fontfile)
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}

	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    80,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	dc.SetFontFace(face)
	dc.SetColor(userColors.TextColor)
	dc.DrawStringWrapped(title, ogImgWidthFloat/2, ogImgHeightFloat*2/6, 0.5, 0.5, 1000, 1.0, gg.AlignCenter)

	// Author
	face, err = opentype.NewFace(f, &opentype.FaceOptions{
		Size:    50,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	dc.SetFontFace(face)
	dc.SetColor(userColors.TextColor)
	dc.DrawStringWrapped(authors, ogImgWidthFloat/2, ogImgHeightFloat*4/6, 0.5, 0.5, 1000, 1.0, gg.AlignCenter)

	// Logo + Site name
	face, err = opentype.NewFace(f, &opentype.FaceOptions{
		Size:    40,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	dc.SetFontFace(face)
	dc.SetColor(userColors.TextColor)

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
	if sitename != "" {
		w, _ := dc.MeasureString(sitename)
		logoRowWidth += int(w)
	}
	if logoimgurl != "" && sitename != "" {
		logoRowWidth += 10.0
	}

	// Draw Logo Images
	logoRowItemAnchor := (ogImgWidth / 2) - (logoRowWidth / 2)
	if logoimgurl != "" {
		logoRowItemAnchor += (logoimg.Bounds().Size().X) / 2
		dc.DrawImageAnchored(logoimg, logoRowItemAnchor, ogImgHeight*5/6, 0.5, 0.5)
		logoRowItemAnchor += (logoimg.Bounds().Size().X) / 2
	}
	if logoimgurl != "" && sitename != "" {
		logoRowItemAnchor += 10.0
	}
	if sitename != "" {
		w, _ := dc.MeasureString(sitename)
		logoRowItemAnchor += int(w / 2)
		dc.DrawStringAnchored(sitename, float64(logoRowItemAnchor), float64(ogImgHeight*5/6), 0.5, 0.3)
	}

	imgfilePath := filepath + ".png"
	dc.SavePNG(imgfilePath)
	return imgfilePath
}
