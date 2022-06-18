package user

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	usageString = `usage %s [file path] [top row] [middle row] [bottom row] [left wing] [top wing] [right wing] [bottom wing]

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
	arglength = 8
)

func CheckIfUsage() {
	if !isArgsLengthOk(arglength) {
		usage()
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, usageString, os.Args[0])
	return
}

func isArgsLengthOk(arglength int) bool {
	if len(os.Args) < arglength + 1 {
		return false
	}
	return true
}

func Input() ([9]int, [12]int) {
	args := os.Args
	topcolorsint := [9]int{}
	sidecolorsint := [12]int{}

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
