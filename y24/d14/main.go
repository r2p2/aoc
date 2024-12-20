package d14

import (
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) string {
	xyRe := `p=([0-9]*),([0-9]*) v=([-?0-9]*),(-?[0-9]*)`
	re := regexp.MustCompile(xyRe)

	width := 101
	height := 103
	middlev := width / 2
	middleh := height / 2

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines[:len(lines)-1] {
		matches := re.FindStringSubmatch(line)
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		dx, _ := strconv.Atoi(matches[3])
		dy, _ := strconv.Atoi(matches[4])

		x += dx * 100
		y += dy * 100
		x = mod(x, width)
		y = mod(y, height)

		if x < middlev && y < middleh {
			q1 += 1
		} else if x > middlev && y < middleh {
			q2 += 1
		} else if x < middlev && y > middleh {
			q3 += 1
		} else if x > middlev && y > middleh {
			q4 += 1
		}
	}
	return strconv.Itoa(q1 * q2 * q3 * q4)
}

func Part2(input string) string {
	xyRe := `p=([0-9]*),([0-9]*) v=([-?0-9]*),(-?[0-9]*)`
	re := regexp.MustCompile(xyRe)

	width := 101
	height := 103

	robots := []robot{}

	lines := strings.Split(input, "\n")
	for _, line := range lines[:len(lines)-1] {
		matches := re.FindStringSubmatch(line)
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		dx, _ := strconv.Atoi(matches[3])
		dy, _ := strconv.Atoi(matches[4])

		robots = append(robots, robot{x, y, dx, dy})
	}

	for i := 1; i < 10_0000; i++ {
		if step(robots, width, height) {
			return strconv.Itoa(i)
		}
	}
	return ">8270<"
}

func mod(a, b int) int {
	return (a%b + b) % b
}

type robot struct {
	x, y, dx, dy int
}

func step(robots []robot, w, h int) bool {
	cache := make([]int, w*h)
	for i := range robots {
		robots[i].x += robots[i].dx
		robots[i].y += robots[i].dy
		robots[i].x = mod(robots[i].x, w)
		robots[i].y = mod(robots[i].y, h)
		cache[robots[i].y*w+robots[i].x] += 1
	}
	for _, c := range cache {
		if c > 1 {
			return false
		}
	}
	return true
}
