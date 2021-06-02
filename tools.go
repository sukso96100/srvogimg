package main

import (
	"crypto/sha256"
	"fmt"
	"image/color"
	"strconv"

	"github.com/fogleman/gg"
)

func parseHexColor(hex string) (color.Color, error) {
	values, err := strconv.ParseUint(hex, 16, 32)

	if err != nil {
		return color.RGBA{0, 0, 0, 0}, err
	}
	red := uint8(values >> 16)
	blue := uint8((values >> 8) & 0xFF)
	green := uint8(values & 0xFF)

	rgb := color.RGBA{red, blue, green, 255}

	return rgb, nil
}

func createBackground(start string, end string) gg.Gradient {
	startRGB, _ := parseHexColor(start)
	endRGB, _ := parseHexColor(end)
	grad := gg.NewLinearGradient(0, 0, 1200, 600)
	grad.AddColorStop(0, startRGB)
	grad.AddColorStop(1, endRGB)
	return grad
}

func getHashedFileName(filename string) string {
	h := sha256.New()
	h.Write([]byte(filename))
	hbytes := h.Sum(nil)
	return fmt.Sprintf("%x", hbytes)
}
