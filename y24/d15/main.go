package d15

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) string {
	area := [][]rune{}
	x := 0
	y := 0
	parseMap := true
	movesBuilder := strings.Builder{}
	lines := strings.Split(input, "\n")
	for _, line := range lines[:len(lines)-1] {
		if line == "" {
			parseMap = false
			continue
		}

		if parseMap {
			area = append(area, []rune(line))
			continue
		}
		movesBuilder.WriteString(line)
	}

	moves := movesBuilder.String()

	for iy := 0; iy < len(area); iy++ {
		for ix := 0; ix < len(area[iy]); ix++ {
			if area[iy][ix] == '@' {
				x = ix
				y = iy
			}
		}
	}

	for _, move := range moves {
		xn := 0
		yn := 0
		if move == '>' {
			xn = x + 1
			yn = y
		} else if move == '<' {
			xn = x - 1
			yn = y
		} else if move == '^' {
			xn = x
			yn = y - 1
		} else if move == 'v' {
			xn = x
			yn = y + 1
		}
		if swap(area, x, y, xn, yn) {
			x = xn
			y = yn
		}
	}

	sum := 0
	for iy := 0; iy < len(area); iy++ {
		for ix := 0; ix < len(area[iy]); ix++ {
			if area[iy][ix] == 'O' {
				sum += iy*100 + ix
			}
		}
	}

	return strconv.Itoa(sum)
}

func Part2(input string) string {
	area := [][]rune{}
	x := 0
	y := 0
	parseMap := true
	movesBuilder := strings.Builder{}
	lines := strings.Split(input, "\n")
	for _, linei := range lines[:len(lines)-1] {
		if linei == "" {
			parseMap = false
			continue
		}

		line := strings.Replace(linei, ".", "..", -1)
		line = strings.Replace(line, "#", "##", -1)
		line = strings.Replace(line, "O", "[]", -1)
		line = strings.Replace(line, "@", "@.", -1)

		if parseMap {
			area = append(area, []rune(line))
			continue
		}
		movesBuilder.WriteString(line)
	}

	moves := movesBuilder.String()

	for iy := 0; iy < len(area); iy++ {
		for ix := 0; ix < len(area[iy]); ix++ {
			if area[iy][ix] == '@' {
				x = ix
				y = iy
			}
		}
	}

	for _, move := range moves {
		if !check(area, x, y, move, false) {
			continue
		}
		if !execute(area, x, y, move, false) {
			continue
		}

		if move == '>' {
			x++
		} else if move == '<' {
			x--
		} else if move == '^' {
			y--
		} else if move == 'v' {
			y++
		}
	}

	sum := 0
	for iy := 0; iy < len(area); iy++ {
		for ix := 0; ix < len(area[iy]); ix++ {
			if area[iy][ix] == '[' {
				sum += iy*100 + ix
			}
		}
	}

	return strconv.Itoa(sum)
}

func check(area [][]rune, x, y int, move rune, pair bool) bool {
	if area[y][x] == '#' {
		return false
	}

	if area[y][x] == '.' {
		return true
	}

	xn := x
	yn := y

	if move == '>' {
		xn++
	} else if move == '<' {
		xn--
	} else if move == '^' {
		yn--
	} else if move == 'v' {
		yn++
	}

	if !check(area, xn, yn, move, false) {
		return false
	}

	if (move == '^' || move == 'v') && !pair {
		if area[y][x] == '[' {
			if !check(area, x+1, y, move, true) {
				return false
			}
		} else if area[y][x] == ']' {
			if !check(area, x-1, y, move, true) {
				return false
			}
		}
	}

	return true
}

func execute(area [][]rune, x, y int, move rune, pair bool) bool {
	if area[y][x] == '#' {
		return false
	}

	if area[y][x] == '.' {
		return true
	}

	xn := x
	yn := y

	if move == '>' {
		xn++
	} else if move == '<' {
		xn--
	} else if move == '^' {
		yn--
	} else if move == 'v' {
		yn++
	}

	if !execute(area, xn, yn, move, false) {
		return false
	}

	if (move == '^' || move == 'v') && !pair {
		if area[y][x] == '[' {
			if !execute(area, x+1, y, move, true) {
				return false
			}
		} else if area[y][x] == ']' {
			if !execute(area, x-1, y, move, true) {
				return false
			}
		}
	}

	area[yn][xn] = area[y][x]
	area[y][x] = '.'

	return true
}

func swap(area [][]rune, xs, ys, xd, yd int) bool {
	if area[yd][xd] == '#' {
		return false
	}

	if area[yd][xd] != '.' {
		if !swap(area, xd, yd, xd+(xd-xs), yd+(yd-ys)) {
			return false
		}
	}
	if (yd-ys) != 0 && area[ys][xs] == '[' {
		if area[yd][xd+1] != '.' {
			if !swap(area, xd+1, yd, xd+1, yd+(yd-ys)) {
				return false
			}
		}
	} else if (yd-ys) != 0 && area[ys][xs] == ']' {
		if area[yd][xd-1] != '.' {
			// fmt.Printf("this:%d,%d, n:%d,%d\n", xs, ys, xd-1, yd+(yd-ys))
			// fmt.Printf("this:%c u/d:%c\n", area[ys][xs], area[yd+(yd-ys)][xd-1])
			if !swap(area, xd-1, yd, xd-1, yd+(yd-ys)) {
				return false
			}
		}
	}

	if area[ys][xs] == '#' {
		return false
	}

	if (yd-ys) != 0 && area[ys][xs] == '[' {
		area[yd][xd+1] = area[ys][xs+1]
		area[ys][xs+1] = '.'
	} else if (yd-ys) != 0 && area[ys][xs] == ']' {
		area[yd][xd-1] = area[ys][xs-1]
		area[ys][xs-1] = '.'
	}
	area[yd][xd] = area[ys][xs]
	area[ys][xs] = '.'
	return true
}

func print(area [][]rune) {
	for _, line := range area {
		for _, c := range line {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}
