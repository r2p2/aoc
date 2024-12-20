package d01

import (
	"sort"
	"strconv"
	"strings"
)

func Part1(input string) string {
	lhs := []int{}
	rhs := []int{}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		l_r := strings.Fields(line)

		l, _ := strconv.Atoi(l_r[0])
		lhs = append(lhs, l)

		r, _ := strconv.Atoi(l_r[1])
		rhs = append(rhs, r)
	}

	sort.Ints(lhs)
	sort.Ints(rhs)

	sum := 0
	for i := range lhs {
		diff := rhs[i] - lhs[i]
		if diff < 0 {
			diff *= -1
		}
		sum += diff
	}

	return strconv.Itoa(sum)
}

func Part2(input string) string {
	lhs := []int{}
	rhs := map[int]int{}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		l_r := strings.Fields(line)

		l, _ := strconv.Atoi(l_r[0])
		lhs = append(lhs, l)

		r, _ := strconv.Atoi(l_r[1])
		rhs[r] += 1
	}

	sum := 0
	for _, l := range lhs {
		sum += l * rhs[l]
	}

	return strconv.Itoa(sum)
}
