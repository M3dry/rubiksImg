package draw

import (
	"errors"
	"image/color"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dkit"
)

func CubeLL(gc draw2d.GraphicContext, topcolors [9]int, sidecolors [12]int, colors [8]color.RGBA) error {
	if twoColors(topcolors, sidecolors) {
		return errors.New("Two or more same colors in one piece")
	}

	gc.MoveTo(0, 0)
	gc.Fill()

	gc.SetFillColor(colors[7])
	draw2dkit.Rectangle(gc, 100, 100, 405, 405)
	gc.Fill()

	//top colors
	x := 0
	y := 100

	for _, value := range topcolors {
		x += 100
		for x > 300 {
			x = 100
			y += 100
		}

		if value < 7 {
			gc.SetFillColor(colors[value])
			draw2dkit.Rectangle(gc, float64(x) + 5, float64(y) + 5, float64(x + 100), float64(y + 100))
			gc.Fill()
		}
	}

	// side colors
	// left wing
	x = 30
	y = 100
	gc.SetFillColor(colors[7])
	draw2dkit.Rectangle(gc, float64(x), float64(y), float64(x + 70), float64(y + 305))
	gc.Fill()
	// ring wing
	x = 400
	y = 100
	gc.SetFillColor(colors[7])
	draw2dkit.Rectangle(gc, float64(x), float64(y), float64(x + 70), float64(y + 305))
	gc.Fill()
	// top wing
	x = 100
	y = 100
	gc.SetFillColor(colors[7])
	draw2dkit.Rectangle(gc, float64(x), float64(y), float64(x + 305), float64(30))
	gc.Fill()
	// bottom wing
	x = 100
	y = 400
	gc.SetFillColor(colors[7])
	draw2dkit.Rectangle(gc, float64(x), float64(y), float64(x + 305), float64(y + 70))
	gc.Fill()

	for index, value := range sidecolors {
		if index < 3 {
			if value < 7 {
				newdex := index + 1

				gc.SetFillColor(colors[value])
				draw2dkit.Rectangle(gc, float64(35), float64(100 + 5 * newdex + 95 * index), float64(100), float64(100 + 5 * newdex + 95 * newdex))
				gc.Fill()
			}
		} else if index < 6 {
			if value < 7 {
				newdex := index - 3 + 1
				x = 100
				y = 100

				gc.SetFillColor(colors[value])
				draw2dkit.Rectangle(gc, float64(100 * newdex + 5), float64(100), float64(100 * newdex + 100), float64(35))
				gc.Fill()
			}
		} else if index < 9 {
			if value < 7 {
				newdex := index - 6 + 1

				gc.SetFillColor(colors[value])
				draw2dkit.Rectangle(gc, float64(405), float64(100 + 5 * newdex + 95 * (index - 6)), float64(465), float64(100 + 5 * newdex + 95 * newdex))
				gc.Fill()
			}
		} else if index < 12 {
			if value < 7 {
				newdex := index - 9 + 1

				gc.SetFillColor(colors[value])
				draw2dkit.Rectangle(gc, float64(100 * newdex + 5), float64(405), float64(100 * newdex + 100), float64(465))
				gc.Fill()
			}
		}
	}

	return nil
}
