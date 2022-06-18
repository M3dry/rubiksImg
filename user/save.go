package user

import (
	"github.com/llgcode/draw2d/draw2dsvg"
)

func Save(file string, svg *draw2dsvg.Svg) {
	draw2dsvg.SaveToSvgFile(file, svg)
}
