package draw

import (
	"testing"
)

type twoColor struct {
	topcolors [9]int
	sidecolors [12]int
	isDuplicate bool
}

func TestTwoColors(t *testing.T) {
	values := []twoColor{
		{ topcolors: [9]int{0,0,0, 0,0,0, 0,0,0}, sidecolors: [12]int{0,0,0, 0,0,0, 0,0,0, 0,0,0}, isDuplicate: true },
		{ topcolors: [9]int{1,1,1, 1,1,1, 1,1,1}, sidecolors: [12]int{2,2,2, 1,1,1, 3,3,3, 4,4,4}, isDuplicate: true },
		{ topcolors: [9]int{1,1,1, 1,1,1, 1,1,1}, sidecolors: [12]int{2,2,2, 3,3,3, 1,1,1, 4,4,4}, isDuplicate: true },
		{ topcolors: [9]int{1,1,1, 1,1,1, 1,1,1}, sidecolors: [12]int{2,2,2, 3,3,3, 4,4,4, 1,1,1}, isDuplicate: true },
		{ topcolors: [9]int{1,1,1, 1,1,1, 1,1,1}, sidecolors: [12]int{5,5,5, 2,2,2, 3,3,3, 4,4,4}, isDuplicate: false },
		{ topcolors: [9]int{2,2,2, 2,2,2, 2,2,2}, sidecolors: [12]int{1,0,0, 1,0,0, 0,0,0, 0,0,0}, isDuplicate: true },
		{ topcolors: [9]int{2,2,2, 2,2,2, 2,2,2}, sidecolors: [12]int{0,0,0, 0,0,1, 1,0,0, 0,0,0}, isDuplicate: true },
		{ topcolors: [9]int{2,2,2, 2,2,2, 2,2,2}, sidecolors: [12]int{0,0,0, 0,0,0, 0,0,1, 0,0,1}, isDuplicate: true },
		{ topcolors: [9]int{2,2,2, 2,2,2, 2,2,2}, sidecolors: [12]int{0,0,1, 0,0,0, 0,0,0, 1,0,0}, isDuplicate: true },
	}

	for _, value := range values {
		result := twoColors(value.topcolors, value.sidecolors)

		if result != value.isDuplicate {
			t.Errorf("twoColors(%d, %d) FAILED, Expected %t, got %t", value.topcolors, value.sidecolors, value.isDuplicate, result)
		}
	}
}
