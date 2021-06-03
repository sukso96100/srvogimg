package draw

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	"io"
	"log"
	"net/http"

	"github.com/sukso96100/srvogimg/res"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const ogImgWidth int = 1200
const ogImgHeight int = 600

func drawBasicOgImage(text string, imgurls []string, startColor string, endColor string, filepath string) string {
	dc := gg.NewContext(1200, 600)

	// Background
	dc.DrawRectangle(0, 0, 1200, 600)
	dc.SetFillStyle(createBackground(startColor, endColor))
	dc.Fill()

	resizedImgs := []image.Image{}
	widthTotal := 0
	padding := 20

	// Load Logo Images
	for _, imgurl := range imgurls {
		if imgurl != "" {
			resp, err := http.Get(imgurl)
			var iconimg []byte
			if err != nil {
				// handle error
				iconimg = res.DefaultLogo
			} else {
				defer resp.Body.Close()
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					iconimg = res.DefaultLogo
				}
				iconimg = body
			}

			img, _, err := image.Decode(bytes.NewReader(iconimg))
			if err != nil {
				log.Fatalln(err)
			}
			m := resize.Resize(0, 200, img, resize.Lanczos3)
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
	dc.DrawStringWrapped(text, 1200/2, 600*3/4, 0.5, 0.5, 1000, 0.8, gg.AlignCenter)
	imgfilePath := filepath + ".png"
	dc.SavePNG(imgfilePath)
	return imgfilePath
}
