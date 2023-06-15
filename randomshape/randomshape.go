package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/jdxyw/generativeart/common"
)

func main() {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(1920, 1080)
	c.SetBackground(common.White)
	c.FillBackground()
	c.SetColorSchema([]color.RGBA{
		{0x5f, 0x9e, 0xa0, 0xff},
		{0xff, 0xe4, 0xc4, 0xff},
		{0x87, 0xce, 0xfa, 0xff},
		{0x55, 0x6b, 0x2f, 0xff},
		{0xff, 0xa0, 0x7a, 0xff},
	})
	c.Draw(arts.NewRandomShape(233))
	c.ToPNG("randomshape3.png")
}
