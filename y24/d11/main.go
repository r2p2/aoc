package d11

import (
	"math"
	"strconv"
	"strings"
)

func Part1(input string) string {
	return strconv.Itoa(partHelper(input, 25))
}

func Part2(input string) string {
	return strconv.Itoa(partHelper(input, 75))
}

type Entry struct {
	iteration int
	value     int
}

func solve(memory []Entry, iterations int) int {
	table := make(map[Entry]int)
	sum := 0
	for _, n := range memory {
		sum += part2(n.value, iterations, &table)
	}
	return sum
}

func part2(n int, iterations int, table *map[Entry]int) int {
	if iterations == 0 {
		return 1
	}

	table_entry, ok := (*table)[Entry{iterations, n}]
	if ok {
		return table_entry
	}

	digits := int(math.Log10(float64(n))) + 1

	result := 0
	if n == 0 {
		result = part2(1, iterations-1, table)
	} else if digits%2 == 0 {
		powDigits := math.Pow(10, float64(digits)-float64(digits/2.0))
		left := int(math.Floor(float64(n) / powDigits))
		right := n % int(powDigits)
		result = part2(int(left), iterations-1, table) + part2(int(right), iterations-1, table)
	} else {
		result = part2(n*2024, iterations-1, table)
	}

	(*table)[Entry{iterations, n}] = result
	return result
}

func partHelper(input string, iterations int) int {
	stones_strings := strings.Split(strings.ReplaceAll(input, "\n", ""), " ")
	stones := []Entry{}
	for _, stone_string := range stones_strings {
		i, _ := strconv.ParseInt(stone_string, 10, 64)
		stones = append(stones, Entry{0, int(i)})
	}
	return solve(stones, iterations)
}
