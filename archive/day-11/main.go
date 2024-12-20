package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	memory1 := []Entry{}
	memory2 := []Entry{}

	input := strings.Split(strings.ReplaceAll(string(bytes), "\n", ""), " ")
	for _, word := range input {
		i, _ := strconv.ParseInt(word, 10, 64)
		memory1 = append(memory1, Entry{0, int(i)})
		memory2 = append(memory2, Entry{0, int(i)})
	}

	fmt.Println("Part 1:", solve(memory1, 25))
	fmt.Println("Part 2:", solve(memory2, 75))
}
