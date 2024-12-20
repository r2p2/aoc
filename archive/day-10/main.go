package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

func step(area []string, width, height, x, y, xold, yold int, niner *[]Pos, visited []Pos) int {
	if x < 0 || y < 0 || y >= height || x >= width {
		return 0
	}

	if slices.Contains(visited, Pos{x, y}) {
		return 0
	}

	if area[y][x] == '.' {
		return 0
	}

	nold, _ := strconv.Atoi(string(area[yold][xold]))
	nnew, _ := strconv.Atoi(string(area[y][x]))
	if (nnew - nold) != 1 {
		return 0
	}

	if area[y][x] == '9' {
		if !slices.Contains(*niner, Pos{x, y}) {
			*niner = append(*niner, Pos{x, y})
		}
		return 1
	}

	sum := 0
	sum += step(area, width, height, x, y-1, x, y, niner, append(visited, Pos{x, y}))
	sum += step(area, width, height, x, y+1, x, y, niner, append(visited, Pos{x, y}))
	sum += step(area, width, height, x-1, y, x, y, niner, append(visited, Pos{x, y}))
	sum += step(area, width, height, x+1, y, x, y, niner, append(visited, Pos{x, y}))
	return sum
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	area := strings.Split(string(bytes), "\n")
	width := len(area[0])
	height := len(area) - 1
	reachable := 0
	paths := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if area[y][x] != '0' {
				continue
			}

			niner := []Pos{}
			paths += step(area, width, height, x, y-1, x, y, &niner, []Pos{Pos{x, y}})
			paths += step(area, width, height, x, y+1, x, y, &niner, []Pos{Pos{x, y}})
			paths += step(area, width, height, x-1, y, x, y, &niner, []Pos{Pos{x, y}})
			paths += step(area, width, height, x+1, y, x, y, &niner, []Pos{Pos{x, y}})
			reachable += len(niner)
		}
	}

	fmt.Println("Part 1:", reachable)
	fmt.Println("Part 2:", paths)
}
