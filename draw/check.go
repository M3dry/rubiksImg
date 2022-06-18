package draw

func twoColors(topcolors [9]int, sidecolors [12]int) bool {
	tmp := 2

	for i, val := range sidecolors {
		if i < 3 {
			if val == topcolors[i * 3] { // topcolors: 0,3,6 sidecolors: 0,1,2
				return true
			}
		} else if i < 6 {
			if val == topcolors[i - 3] { // topcolors: 0,1,2 sidecolors: 3,4,5
				return true
			}
		} else if i < 9 {
			if val ==  topcolors[tmp] { // topcolors: 2,5,8 sidecolors: 6,7,8
				return true
			}

			tmp += 3
		} else if i < 12 { // topcolors: 6,7,8 sidecolors: 9,10,11
			if val ==  topcolors[i - 3] {
				return true
			}
		}
	}

	if sidecolors[0] == sidecolors[3] && sidecolors[0] == 1 {
		return true
	}

	if sidecolors[5] == sidecolors[6] && sidecolors[5] == 1 {
		return true
	}

	if sidecolors[8] == sidecolors[11] && sidecolors[8] == 1 {
		return true
	}

	if sidecolors[2] == sidecolors[9] && sidecolors[2] == 1 {
		return true
	}

	return false
}
