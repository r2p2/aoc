package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Pos struct {
	x, y int
}
type PosD struct {
	x, y int
	d    rune
}

func part1(area []string, w, h int, pos Pos, d rune) []Pos {
	visited := make([]Pos, 0, 1024)
	// visited := []Pos{}
	for {
		if !slices.Contains(visited, pos) {
			visited = append(visited, pos)
		}

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
		idx := pos.y*130*4 + pos.x*4 + dir(d)
		if visited[idx] {
			return true
		}
		visited[idx] = true

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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("where is the file?")
		return
	}

	path := os.Args[1]

	input, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		return
	}
	area := strings.Split(string(input), "\n")
	height := len(area) - 1
	width := len(area[0])
	fmt.Println(height, width)
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
	fmt.Println("Part 1:", len(visited))

	c := make(chan bool, 100)
	for i := 0; i < len(visited); i++ {
		go func(c chan bool) {
			if part2(area, width, height, pos, d, visited[i].x, visited[i].y) {
				c <- true
			} else {
				c <- false
			}
		}(c)
	}

	sum := 0
	for i := 0; i < len(visited); i++ {
		if <-c {
			sum += 1
		}
	}

	fmt.Println("Part 2:", sum)
}
