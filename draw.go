package main

import (
	"bytes"
	"image"
	"image/color"
	"io"
	"log"
	"net/http"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func drawOgImage(text string, imgurl string, startColor string, endColor string, filepath string) string {
	dc := gg.NewContext(1200, 600)

	// Background
	dc.DrawRectangle(0, 0, 1200, 600)
	dc.SetFillStyle(createBackground(startColor, endColor))
	dc.Fill()

	// Image
	if imgurl != "" {
		resp, err := http.Get(imgurl)
		var iconimg []byte
		if err != nil {
			// handle error
			iconimg = defaultImg
		} else {
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				iconimg = defaultImg
			}
			iconimg = body
		}

		img, _, err := image.Decode(bytes.NewReader(iconimg))
		if err != nil {
			log.Fatalln(err)
		}
		m := resize.Resize(0, 200, img, resize.Lanczos3)
		dc.DrawImageAnchored(m, 1200/2, 600*1/3, 0.5, 0.5)
	}

	// Text
	f, err := opentype.Parse(fontfile)
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}

	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    80,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	dc.SetFontFace(face)
	dc.SetColor(color.White)
	dc.DrawStringWrapped(text, 1200/2, 600*3/4, 0.5, 0.5, 1000, 0.8, gg.AlignCenter)
	dc.SavePNG(filepath)
	return filepath
}
