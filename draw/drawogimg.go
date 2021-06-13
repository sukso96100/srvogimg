package draw

import (
	_ "embed"
	"image"
	"image/color"
	"log"

	"github.com/sukso96100/srvogimg/res"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func drawBasicOgImage(text string, imgurls []string, startColor string, endColor string, filepath string) string {

	dc := gg.NewContext(ogImgWidth, ogImgHeight)

	// Background
	dc.DrawRectangle(0, 0, ogImgWidthFloat, ogImgHeightFloat)
	dc.SetFillStyle(NewGradientBackground(startColor, endColor, 255))
	dc.Fill()

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
	dc.SetColor(color.White)
	dc.DrawStringWrapped(text, ogImgWidthFloat/2, ogImgHeightFloat*3/4, 0.5, 0.5, 1000, 0.8, gg.AlignCenter)

	imgfilePath := filepath + ".png"
	dc.SavePNG(imgfilePath)
	return imgfilePath
}
