package main

import (
	"aoc/y24/d01"
	"aoc/y24/d02"
	"aoc/y24/d03"
	"aoc/y24/d06"
	"aoc/y24/d11"
	"aoc/y24/d12"
	"aoc/y24/d13"
	"aoc/y24/d14"
	"aoc/y24/d15"
	"aoc/y24/d16"
	"aoc/y24/d17"
	"aoc/y24/d18"
	"aoc/y24/d19"
	"aoc/y24/d20"
	"aoc/y24/d21"
	"aoc/y24/d22"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"
)

var challenges = []Challenge{
	{24, 1, d01.Part1, d01.Part2},
	{24, 2, d02.Part1, d02.Part2},
	{24, 3, d03.Part1, d03.Part2},
	{24, 6, d06.Part1, d06.Part2},
	{24, 11, d11.Part1, d11.Part2},
	{24, 12, d12.Part1, d12.Part2},
	{24, 13, d13.Part1, d13.Part2},
	{24, 14, d14.Part1, d14.Part2},
	{24, 15, d15.Part1, d15.Part2},
	{24, 16, d16.Part1, d16.Part2},
	{24, 17, d17.Part1, d17.Part2},
	{24, 18, d18.Part1, d18.Part2},
	{24, 19, d19.Part1, d19.Part2},
	{24, 20, d20.Part1, d20.Part2},
	{24, 21, d21.Part1, d21.Part2},
	{24, 22, d22.Part1, d22.Part2},
}

type Challenge struct {
	year, day    int
	part1, part2 func(string) string
}

func run(c Challenge, part int, input string, solution *string) {

	result := ""

	start := time.Now()
	if part == 1 {
		result = c.part1(input)
	} else {
		result = c.part2(input)
	}
	delta := time.Now().Sub(start)

	status := "????"
	comparison := ""
	if solution != nil && *solution == result {
		status = " OK "
		comparison = fmt.Sprintf(" == %s", *solution)
	} else if solution != nil && *solution != result {
		status = "FAIL"
		comparison = fmt.Sprintf(" != %s", *solution)
	}

	fmt.Printf("[%s] %d.%02d:%d (%13v): %20s%s \n", status, c.year, c.day, part, delta, result, comparison)
}

func loadInput(year, day int) (string, error) {
	bytes, err := os.ReadFile(fmt.Sprintf("inputs/%d/%d.txt", year, day))
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func main() {
	profcpu := flag.String("profcpu", "", "profile cpu")
	profmem := flag.String("profmem", "", "profile memory")
	year := flag.Int("year", 0, "year or 0")
	day := flag.Int("day", 0, "day or 0")
	part := flag.Int("part", 0, "part or 0")
	inputFile := flag.String("input", "", "alternative input file (requires year and date)")
	flag.Parse()

	if *profcpu != "" {
		f, err := os.Create(*profcpu)
		if err != nil {
			log.Fatal("could not create cpu profile: ", err)
		}

		defer f.Close()

		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start cpu profiling: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if *profmem != "" {
		f, err := os.Create(*profmem)
		if err != nil {
			log.Fatal("could not create mem profile: ", err)
		}

		defer f.Close()

		runtime.GC()

		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write mem profile: ", err)
		}
	}

	if len(*inputFile) > 0 && *year == 0 {
		fmt.Println("if input ist provdied, year is required")
		return
	}

	if len(*inputFile) > 0 && *day == 0 {
		fmt.Println("if input ist provdied, day is required")
		return
	}

	for _, c := range challenges {
		if *year != 0 && *year != c.year {
			continue
		}

		if *day != 0 && *day != c.day {
			continue
		}

		var inputPath string
		solutionPath := ""
		var input string

		if *inputFile != "" {
			inputPath = *inputFile
		} else {
			inputPath = fmt.Sprintf("data/%d/%d.txt", c.year, c.day)
			solutionPath = fmt.Sprintf("data/%d/%d_solution.txt", c.year, c.day)
		}

		bytes, err := os.ReadFile(inputPath)
		if err != nil {
			fmt.Println("unable to load test data", inputPath, err)
			continue
		}
		input = string(bytes)

		solution := []*string{}
		if len(solutionPath) != 0 {
			bytes, err := os.ReadFile(solutionPath)
			if err != nil {
				empty := ""
				solution = []*string{&empty, &empty}
			} else {
				for _, line := range strings.Split(string(bytes), "\n") {
					solution = append(solution, &line)
				}
			}
		}

		for len(solution) < 2 {
			solution = append(solution, nil)
		}

		if *part == 0 || *part == 1 {
			run(c, 1, input, solution[0])
		}
		if *part == 0 || *part == 2 {
			run(c, 2, input, solution[1])
		}
	}
}
