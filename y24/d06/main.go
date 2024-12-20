package d06

import (
	"strconv"
	"strings"
)

func Part1(input string) string {
	area := strings.Split(string(input), "\n")
	height := len(area) - 1
	width := len(area[0])
	pos := Pos{
		x: 0,
		y: 0,
	}
	d := '^'
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if area[y][x] == '^' {
				pos.x = x
				pos.y = y
				goto initComplete
			}
		}
	}
initComplete:
	sum := 0
	for _, cell := range part1(area, width, height, pos, d) {
		sum += cell
	}
	return strconv.Itoa(sum)
}

func Part2(input string) string {
	area := strings.Split(string(input), "\n")
	height := len(area) - 1
	width := len(area[0])
	pos := Pos{
		x: 0,
		y: 0,
	}
	d := '^'
	for y := 0; y < len(area); y++ {
		for x := 0; x < len(area[y]); x++ {
			if area[y][x] == '^' {
				pos.x = x
				pos.y = y
			}
		}
	}

	visited := part1(area, width, height, pos, d)

	c := make(chan bool, 100)
	for i, v := range visited {
		if v == 0 {
			continue
		}
		go func(c chan bool) {
			if part2(area, width, height, pos, d, i%width, i/width) {
				c <- true
			} else {
				c <- false
			}
		}(c)
	}

	sum := 0
	for _, v := range visited {
		if v == 0 {
			continue
		}
		if <-c {
			sum += 1
		}
	}

	return strconv.Itoa(sum)
}

type Pos struct {
	x, y int
}
type PosD struct {
	x, y int
	d    rune
}

func part1(area []string, w, h int, pos Pos, d rune) []int {
	visited := make([]int, w*h)
	for {
		visited[pos.y*w+pos.x] = 1

		pos_new := pos
		switch d {
		case '^':
			pos_new.y--
		case '>':
			pos_new.x++
		case 'v':
			pos_new.y++
		case '<':
			pos_new.x--
		}

		out_of_bounds := pos_new.x < 0 || pos_new.y < 0 || pos_new.x >= w || pos_new.y >= h
		if out_of_bounds {
			return visited
		}

		if area[pos_new.y][pos_new.x] == '#' {
			switch d {
			case '^':
				d = '>'
			case '>':
				d = 'v'
			case 'v':
				d = '<'
			case '<':
				d = '^'
			}
			continue
		}
		pos = pos_new
	}
}

func dir(d rune) int {
	switch d {
	case '^':
		return 0
	case '>':
		return 1
	case 'v':
		return 2
	case '<':
		return 3
	}
	return -1
}

func part2(area []string, w, h int, pos Pos, d rune, block_x, block_y int) bool {
	visited := make([]bool, 130*130*4)
	for {
		pos_new := pos
		switch d {
		case '^':
			pos_new.y--
		case '>':
			pos_new.x++
		case 'v':
			pos_new.y++
		case '<':
			pos_new.x--
		}

		out_of_bounds := pos_new.x < 0 || pos_new.y < 0 || pos_new.x >= w || pos_new.y >= h
		if out_of_bounds {
			return false
		}

		if area[pos_new.y][pos_new.x] == '#' || (block_x == pos_new.x && block_y == pos_new.y) {
			idx := pos.y*130*4 + pos.x*4 + dir(d)
			if visited[idx] {
				return true
			}
			visited[idx] = true

			switch d {
			case '^':
				d = '>'
			case '>':
				d = 'v'
			case 'v':
				d = '<'
			case '<':
				d = '^'
			}
			continue
		}
		pos = pos_new
	}
}
