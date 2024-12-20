package d12

import (
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func Part1(input string) string {
	area := strings.Split(input, "\n")
	height := len(area) - 1
	width := len(area[0])
	cache := make([]bool, width*height)

	sum := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if cache[y*width+x] {
				continue
			}
			sizeArea, sizeBorder := floodFill(area, width, height, Point{x, y}, area[y][x], &cache)
			sum += sizeArea * sizeBorder
		}
	}
	return strconv.Itoa(sum)
}

func Part2(input string) string {
	area := strings.Split(input, "\n")
	height := len(area) - 1
	width := len(area[0])
	cache := make([]bool, width*height)

	sum := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if cache[y*width+x] {
				continue
			}
			cacheLocal := make([]bool, width*height)
			sizeArea, _ := floodFill(area, width, height, Point{x, y}, area[y][x], &cacheLocal)
			for i, p := range cacheLocal {
				if p {
					cache[i] = p
				}
			}
			sum += sizeArea * countCorners(cacheLocal, width, height)
		}
	}
	return strconv.Itoa(sum)
}

func countCorners(cache []bool, width, height int) int {
	corners := 0
	for i, p := range cache {
		if !p {
			continue
		}

		y := i / width
		x := i % width

		t := 0
		d := 0
		l := 0
		r := 0
		tl := 0
		tr := 0
		dl := 0
		dr := 0
		if y > 0 && cache[(y-1)*width+x] {
			t = 1
		}
		if y < height-1 && cache[(y+1)*width+x] {
			d = 1
		}
		if x > 0 && cache[y*width+x-1] {
			l = 1
		}
		if x < width-1 && cache[y*width+x+1] {
			r = 1
		}
		if x > 0 && y > 0 && cache[(y-1)*width+x-1] {
			tl = 1
		}
		if x < width-1 && y > 0 && cache[(y-1)*width+x+1] {
			tr = 1
		}
		if x > 0 && y < height-1 && cache[(y+1)*width+x-1] {
			dl = 1
		}
		if x < width-1 && y < height-1 && cache[(y+1)*width+x+1] {
			dr = 1
		}

		if t == 0 && l == 0 {
			corners += 1
		}
		if t == 0 && r == 0 {
			corners += 1
		}
		if d == 0 && l == 0 {
			corners += 1
		}
		if d == 0 && r == 0 {
			corners += 1
		}

		if l == 1 && d == 1 && dl == 0 {
			corners += 1
		}
		if l == 1 && t == 1 && tl == 0 {
			corners += 1
		}
		if t == 1 && r == 1 && tr == 0 {
			corners += 1
		}
		if r == 1 && d == 1 && dr == 0 {
			corners += 1
		}
	}
	return corners
}

func floodFill(area []string, width, height int, point Point, r byte, cache *[]bool) (int, int) {
	sizeArea := 1
	sizeBorder := 0

	if point.x < 0 || point.y < 0 {
		return 0, 0
	}

	if point.x >= width || point.y >= height {
		return 0, 0
	}

	if area[point.y][point.x] != r {
		return 0, 0
	}

	if (*cache)[point.y*width+point.x] {
		return 0, 0
	}

	{ // calc border
		if point.x == width-1 || area[point.y][point.x+1] != r {
			sizeBorder += 1
		}
		if point.x == 0 || area[point.y][point.x-1] != r {
			sizeBorder += 1
		}
		if point.y == height-1 || area[point.y+1][point.x] != r {
			sizeBorder += 1
		}
		if point.y == 0 || area[point.y-1][point.x] != r {
			sizeBorder += 1
		}
	}

	(*cache)[point.y*width+point.x] = true

	{
		sa, sb := floodFill(area, width, height, Point{point.x, point.y + 1}, r, cache)
		sizeArea += sa
		sizeBorder += sb
	}
	{
		sa, sb := floodFill(area, width, height, Point{point.x, point.y - 1}, r, cache)
		sizeArea += sa
		sizeBorder += sb
	}
	{
		sa, sb := floodFill(area, width, height, Point{point.x + 1, point.y}, r, cache)
		sizeArea += sa
		sizeBorder += sb
	}
	{
		sa, sb := floodFill(area, width, height, Point{point.x - 1, point.y}, r, cache)
		sizeArea += sa
		sizeBorder += sb
	}

	return sizeArea, sizeBorder
}
