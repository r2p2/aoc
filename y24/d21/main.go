package d21

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var cache map[int]map[string]int

var (
	dirKeyPadPaths = map[task][]string{
		{'A', '^'}: {"<"},
		{'A', '>'}: {"v"},
		{'A', 'v'}: {"v<", "<v"},
		{'A', '<'}: {"v<<", "<v<"},

		{'^', 'A'}: {">"},
		{'^', '>'}: {">v", "v>"},
		{'^', 'v'}: {"v"},
		{'^', '<'}: {"v<"},

		{'>', 'A'}: {"^"},
		{'>', '^'}: {"^<", "<^"},
		{'>', 'v'}: {"<"},
		{'>', '<'}: {"<<"},

		{'v', '^'}: {"^"},
		{'v', '>'}: {">"},
		{'v', 'A'}: {">^", "^>"},
		{'v', '<'}: {"<"},

		{'<', '^'}: {">^"},
		{'<', '>'}: {">>"},
		{'<', 'v'}: {">"},
		{'<', 'A'}: {">>^", ">^>"},
	}
	numKeyPadPaths = map[task]paths{}
)

func Part1(input string) string {
	cache = map[int]map[string]int{}
	initNumPadPaths()

	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		code, _ := strconv.Atoi(line[:len(line)-1])
		presses := solveMainRobot(line, 2)
		sum += code * presses
	}

	return strconv.Itoa(sum)
}

func Part2(input string) string {
	cache = map[int]map[string]int{}
	initNumPadPaths()

	sum := 0

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		code, _ := strconv.Atoi(line[:len(line)-1])
		presses := solveMainRobot(line, 25)
		sum += code * presses
	}

	return strconv.Itoa(sum)
}

func solveMainRobot(code string, nrAssistants int) int {
	var (
		countMyMoves int
	)
	curr := 'A'
	for _, c := range code {
		paths := numKeyPadPaths[task{key(curr), key(c)}]
		numPadActivations := []int{}
		for _, p := range paths {
			numPadActivations = append(numPadActivations, solveAssistantRobot(nrAssistants, string(p)))
		}

		// TODO get len of p for shortest numpadactivations
		countMyMoves += slices.Min(numPadActivations)

		curr = c
	}
	return countMyMoves
}

func solveAssistantRobot(level int, code string) int {
	code = code + "A"
	if level == 0 {
		// this is me
		return len(code)
	}

	if _, ok := cache[level]; !ok {
		cache[level] = make(map[string]int)
	}
	cache_lvl := cache[level]
	if value, ok := cache_lvl[code]; ok {
		return value
	}

	n := 0
	curr := 'A'
	for _, c := range code {
		if curr == c {
			n += 1 // A
			continue
		}
		paths := dirKeyPadPaths[task{key(curr), key(c)}]
		if len(paths) == 0 {
			panic(fmt.Sprintf("missing entry for num key pad %c -> %c", curr, c))
		}

		klicks := []int{}
		for i := range paths {
			klicks = append(klicks, solveAssistantRobot(level-1, paths[i]))
		}

		n += slices.Min(klicks)
		curr = c
	}

	cache_lvl[code] = n

	return n
}

type pos struct {
	x, y int
}

type key byte
type task struct {
	src, dst key
}

type path string
type paths []path

func initNumPadPaths() {
	keys := []key{'A', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for _, s := range keys {
		for _, d := range keys {
			if s == d {
				continue
			}

			numKeyPadPaths[task{s, d}] = numPadPath(byte(s), byte(d))
		}

	}
}

func numPadPath(src, dst byte) paths {
	nextMap := map[byte][]struct {
		number byte
		move   byte
	}{
		'A': {{'0', '<'}, {'3', '^'}},
		'0': {{'A', '>'}, {'2', '^'}},
		'1': {{'2', '>'}, {'4', '^'}},
		'2': {{'1', '<'}, {'0', 'v'}, {'3', '>'}, {'5', '^'}},
		'3': {{'A', 'v'}, {'2', '<'}, {'6', '^'}},
		'4': {{'1', 'v'}, {'5', '>'}, {'7', '^'}},
		'5': {{'4', '<'}, {'2', 'v'}, {'6', '>'}, {'8', '^'}},
		'6': {{'3', 'v'}, {'5', '<'}, {'9', '^'}},
		'7': {{'8', '>'}, {'4', 'v'}},
		'8': {{'7', '<'}, {'5', 'v'}, {'9', '>'}},
		'9': {{'6', 'v'}, {'8', '<'}},
	}
	pathsToTarget := paths{}

	type Node struct {
		pos      byte
		path     path
		traveled []byte
	}
	queue := []Node{{pos: src, path: "", traveled: []byte{}}}
	for len(queue) > 0 {
		var node Node
		node, queue = queue[0], queue[1:]

		if node.pos == dst {
			pathsToTarget = append(pathsToTarget, node.path)
			continue
		}

		for _, next := range nextMap[node.pos] {
			if slices.Contains(node.traveled, next.number) {
				continue
			}

			newTraveled := make([]byte, len(node.traveled))
			copy(newTraveled, node.traveled)
			newTraveled = append(newTraveled, next.number)

			queue = append(
				queue,
				Node{next.number, path(string(node.path) + string(next.move)), newTraveled},
			) // lol thanks golang
		}
	}

	slices.SortFunc(pathsToTarget, func(a, b path) int {
		return cmp.Compare(len(a), len(b))
	})
	shortestPathsToTarget := slices.DeleteFunc(pathsToTarget, func(path path) bool {
		return len(path) > len(pathsToTarget[0])
	})

	return shortestPathsToTarget
}
