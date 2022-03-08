package main

import (
	"image/color"
	"log"

	"time"

	tjpgd "github.com/sago35/go-tjpgd"
	"github.com/sago35/tinydisplay/examples/initdisplay"
)

var (
	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	green = color.RGBA{0, 255, 0, 255}
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	display := initdisplay.InitDisplay()

	tjpgd.SetCallback(func(left, top, right, bottom uint16, buf []uint16) {
		w := int16(right - left + 1)
		h := int16(bottom - top + 1)
		display.DrawRGBBitmap(int16(left), int16(top), buf, w, h)
	})

	var b []byte
	for i := 0; i < 1000; i++ {
		switch i % 2 {
		case 0:
			b = []byte(img1)
			tjpgd.DecodeFromBytes(b, tjpgd.ScaleNone)
		default:
			b = []byte(img2)
			tjpgd.DecodeFromBytes(b, tjpgd.ScaleNone)
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}
