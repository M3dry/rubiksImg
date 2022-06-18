package main

import (
	"fmt"
	"os"

	"image/color"

	"github.com/llgcode/draw2d/draw2dsvg"

	"github.com/M3dry/rubiksImg/draw"
	"github.com/M3dry/rubiksImg/user"
)

func main() {
	user.CheckIfUsage()
	
	svg := draw2dsvg.NewSvg()
	gc := draw2dsvg.NewGraphicContext(svg)
	
	gc.Save()

	cubecolors := [8]color.RGBA{{255, 255, 255, 0xFF}, {255, 213, 0,  0xFF}, // white[0], yellow[1]
		                        {0,   70,  173, 0xFF}, {0,   155, 72, 0xFF}, // blue[2], green[3]
		                        {183, 18,  52,  0xFF}, {255, 88,  0,  0xFF}, // red[4], orange[5]
	                            {50, 50, 50, 0xFF},    {29,  29,  29, 0xFF}} // empty side[6], base[7]
	topcolorsint, sidecolorsint := user.Input() // get user inputted colors into vars

	if err := draw.CubeLL(gc, topcolorsint, sidecolorsint, cubecolors); err != nil { // Draw last layer 
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}
	
	// save svg to a file defined by user
	gc.Restore()
	user.Save(os.Args[1], svg)
}
