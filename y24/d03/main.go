package d03

import (
	"regexp"
	"strconv"
)

func Part1(input string) string {
	sum := 0
	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	for _, sm := range re.FindAllStringSubmatch(input, -1) {
		l, _ := strconv.Atoi(sm[1])
		r, _ := strconv.Atoi(sm[2])
		sum += l * r
	}

	return strconv.Itoa(sum)
}

func Part2(input string) string {
	sum := 0
	on := true
	re := regexp.MustCompile(`(mul\(([0-9]{1,3}),([0-9]{1,3})\)|don\'t\(\)|do\(\))`)
	for _, sm := range re.FindAllStringSubmatch(input, -1) {
		if sm[0] == "do()" {
			on = true
			continue
		} else if sm[0] == "don't()" {
			on = false
			continue
		}

		if !on {
			continue
		}

		l, _ := strconv.Atoi(sm[2])
		r, _ := strconv.Atoi(sm[3])
		sum += l * r
	}

	return strconv.Itoa(sum)
}
