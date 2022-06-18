// Copyright 2010 The draw2d Authors. All rights reserved.
// created: 21/11/2010 by Laurent Le Goff

// Package gopher2 draws a gopher avatar based on a svg of:
// https://github.com/golang-samples/gopher-vector/
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"image/color"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dkit"
	"github.com/llgcode/draw2d/draw2dsvg"
)

const usage = `usage %s [file path] [top row] [middle row] [bottom row] [left wing] [top wing] [right wing] [bottom wing]

Rows:
	3 numbers from 0 to 6 separated by a colon
Wings:
	3 numbers from 0 to 6 separated by a colon
Colors:
	0: white
	1: yello
	2: blue
	3: green
	4: red
	5: orange
	6: empty
`

func main() {
	args := os.Args
	if len(args) < 9 { // check if user inputted enough arguments, if not print error msg and exit
		fmt.Fprintf(os.Stderr, usage, args[0])
		return
	}
	
	svg := draw2dsvg.NewSvg()
	gc := draw2dsvg.NewGraphicContext(svg)
	cubecolors := [8]color.RGBA{{255, 255, 255, 0xFF}, {255, 213, 0,  0xFF}, // white[0], yellow[1]
		                        {0,   70,  173, 0xFF}, {0,   155, 72, 0xFF}, // blue[2], green[3]
		                        {183, 18,  52,  0xFF}, {255, 88,  0,  0xFF}, // red[4], orange[5]
	                            {50, 50, 50, 0xFF},    {29,  29,  29, 0xFF}} // empty side[6], base[7]
	gc.Save()

	topcolorsint, sidecolorsint := userInput(args) // get user inputted colors into vars
	drawCubeLL(gc, topcolorsint, sidecolorsint, cubecolors)
	
	// save svg to a file defined by user
	gc.Restore()
	draw2dsvg.SaveToSvgFile(args[1], svg)
}

func userInput(args []string) ([9]int, [12]int) {
	topcolorsint := [9]int{}
	sidecolorsint := [12]int{}
	si

	for i := 0; i < 4; i++ {
		if i < 3 {
			for index, value := range strings.Split(args[i+2], ",") {
				if index > 3 {
					break
				}

				topcolorsint[index + i * 3], _ = strconv.Atoi(value)
			}
		}

		for index, value := range strings.Split(args[i + 5], ",") {
			if index > 3 {
				break
			}

			sidecolorsint[index + i * 3], _ = strconv.Atoi(value)
		}
	}

	return topcolorsint, sidecolorsint
}

func drawCubeLL(gc draw2d.GraphicContext, topcolors [9]int, sidecolors [12]int, colors [8]color.RGBA) {
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
}
