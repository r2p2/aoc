package d02

import (
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) string {
	sum := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines[:len(lines)-1] {
		sum += isSafe(strings.Fields(line))
	}

	return strconv.Itoa(sum)
}

func Part2(input string) string {
	sum := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines[:len(lines)-1] {
		reports := strings.Fields(line)
		for i := range reports {
			if isSafe2(reports, i) {
				sum += 1
				break
			}
		}
	}

	return strconv.Itoa(sum)
}

func isSafe(reports []string) int {
	deltas := []int{}
	for i := 0; i < len(reports)-1; i++ {
		l, _ := strconv.Atoi(reports[i])
		r, _ := strconv.Atoi(reports[i+1])
		d := r - l
		if d == 0 {
			return 0
		}
		deltas = append(deltas, d)
	}

	min := slices.Min(deltas)
	max := slices.Max(deltas)

	if min < 0 && max > 0 {
		return 0
	}

	if min < -3 || max > 3 {
		return 0
	}

	return 1
}

func isSafe2(reports []string, ignoreI int) bool {
	deltas := []int{}
	i := 1
	if ignoreI == 0 {
		i = 2
	}
	for ; i < len(reports); i++ {
		prevI := i - 1
		if i == ignoreI {
			if i == len(reports)-1 {
				break
			}
			i++
		}

		l, _ := strconv.Atoi(reports[prevI])
		r, _ := strconv.Atoi(reports[i])
		d := r - l
		if d == 0 {
			return false
		}
		deltas = append(deltas, d)
	}

	min := slices.Min(deltas)
	max := slices.Max(deltas)

	if min < 0 && max > 0 {
		return false
	}

	if min < -3 || max > 3 {
		return false
	}

	return true
}
