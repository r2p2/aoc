package d19

import (
	"strconv"
	"strings"
)

func Part1(input string) string {
	lines := strings.Split(input, "\n")
	towelsLine := lines[0]
	towels := strings.Split(towelsLine, ", ")

	matchCount := 0
	for i := 2; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}

		cache := map[string]int{}
		if matches(lines[i], towels, &cache) > 0 {
			matchCount++
		}
	}

	return strconv.Itoa(matchCount)
}

func Part2(input string) string {
	lines := strings.Split(input, "\n")
	towelsLine := lines[0]
	towels := strings.Split(towelsLine, ", ")

	combinations := 0
	for i := 2; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}

		cache := map[string]int{}
		combinations += matches(lines[i], towels, &cache)
	}

	return strconv.Itoa(combinations)
}

func matches(pattern string, towels []string, cache *map[string]int) int {
	for _, towel := range towels {
		if !strings.HasPrefix(pattern, towel) {
			continue
		}

		if len(towel) > len(pattern) {
			continue
		}

		if len(towel) == len(pattern) {
			if _, ok := (*cache)[pattern]; !ok {
				(*cache)[pattern] = 0
			}
			(*cache)[pattern]++
			continue
		}

		rest := pattern[len(towel):]
		if matchesRest, ok := (*cache)[rest]; ok {
			(*cache)[pattern] += matchesRest
			continue
		}

		(*cache)[pattern] += matches(rest, towels, cache)
	}
	return (*cache)[pattern]
}
