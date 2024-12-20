package d16

import (
	"cmp"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) string {
	area := strings.Split(input, "\n")
	visited := map[pos]int{}
	for y := range area {
		for x := range area[y] {
			visited[pos{x, y, North, 0}] = math.MaxInt
			visited[pos{x, y, East, 0}] = math.MaxInt
			visited[pos{x, y, South, 0}] = math.MaxInt
			visited[pos{x, y, West, 0}] = math.MaxInt
		}
	}
	for y := range area {
		for x := range area[0] {
			if area[y][x] == 'S' {
				ex, ey, _ := bfs(area, pos{x, y, East, 0}, &visited)
				results := []int{
					visited[pos{ex, ey, North, 0}],
					visited[pos{ex, ey, East, 0}],
					visited[pos{ex, ey, South, 0}],
					visited[pos{ex, ey, West, 0}],
				}
				return strconv.Itoa(slices.Min(results))
			}
		}
	}

	return "???"
}

func Part2(input string) string {
	area := strings.Split(input, "\n")
	visited := map[pos]int{}
	for y := range area {
		for x := range area[y] {
			visited[pos{x, y, North, 0}] = math.MaxInt
			visited[pos{x, y, East, 0}] = math.MaxInt
			visited[pos{x, y, South, 0}] = math.MaxInt
			visited[pos{x, y, West, 0}] = math.MaxInt
		}
	}
	for y := range area {
		for x := range area[0] {
			if area[y][x] == 'S' {
				_, _, paths := bfs(area, pos{x, y, East, 0}, &visited)
				slices.SortFunc(paths, func(l, r []pos) int {
					return cmp.Compare(l[len(l)-1].cost, r[len(r)-1].cost)
				})
				uniquePoints := map[point]bool{}
				min := paths[0][len(paths[0])-1].cost
				for i := range paths {
					if paths[i][len(paths[i])-1].cost != min {
						break
					}

					for _, p := range paths[i] {
						uniquePoints[point{p.x, p.y}] = true
					}
				}
				return strconv.Itoa(len(uniquePoints))
			}
		}
	}

	return "???"
}

type point struct {
	x, y int
}

type pos struct {
	x, y int
	dir  dir
	cost int
}

type dir int

const (
	East dir = iota
	South
	West
	North
)

func left(dir dir) dir {
	switch dir {
	case East:
		return North
	case South:
		return East
	case West:
		return South
	case North:
		return West
	}
	return North
}

func right(dir dir) dir {
	switch dir {
	case East:
		return South
	case South:
		return West
	case West:
		return North
	case North:
		return East
	}
	return North
}

func cost(old, new dir) int {
	if old == new {
		return 0
	}
	return 1000
}

func findPaths(costs map[pos]int, curr, end point) int {
	path := []point{curr}
	queue := []point{curr}
	for len(queue) > 0 {
		var e point
		e, queue = queue[0], queue[1:]
		if e.x == end.x && e.y == end.y {
			continue
		}

		dirs := []dir{North, South, East, West}
		slices.SortFunc(dirs, func(a, b dir) int {
			l := costs[pos{e.x, e.y, a, 0}]
			r := costs[pos{e.x, e.y, b, 0}]
			return cmp.Compare(l, r)
		})
		minCoast := costs[pos{e.x, e.y, dirs[0], 0}]

		println("check", e.x, e.y)
		for i := range dirs {
			println(" ", i, costs[pos{e.x, e.y, dirs[i], 0}])
			if !(minCoast == costs[pos{e.x, e.y, dirs[i], 0}] || minCoast == (costs[pos{e.x, e.y, dirs[i], 0}]-1000)) {
				break
			}

			next := e
			switch dirs[i] {
			case North:
				next.y++
			case East:
				next.x--
			case South:
				next.y--
			case West:
				next.x++
			}

			if slices.Contains(path, next) {
				continue
			}

			path = append(path, next)
			queue = append(queue, next)
		}
	}

	return len(path)
}

func bfs(area []string, root pos, visited *map[pos]int) (ex, ey int, pathsToTarget [][]pos) {
	pathsToTarget = [][]pos{}

	type Node struct {
		pos  pos
		path []pos
	}
	(*visited)[root] = 0
	queue := []Node{{pos: root, path: []pos{root}}}
	for len(queue) > 0 {
		var node Node
		node, queue = queue[0], queue[1:]

		if area[node.pos.y][node.pos.x] == 'E' {
			ex = node.pos.x
			ey = node.pos.y
			pathsToTarget = append(pathsToTarget, node.path)
		}

		for _, dir := range []dir{node.pos.dir, left(node.pos.dir), right(node.pos.dir)} {
			next := pos{
				x:    node.pos.x,
				y:    node.pos.y,
				dir:  dir,
				cost: node.pos.cost + 1 + cost(node.pos.dir, dir),
			}

			switch dir {
			case North:
				next.y--
			case East:
				next.x++
			case South:
				next.y++
			case West:
				next.x--
			}

			if area[next.y][next.x] == '#' {
				continue
			}

			if cost, ok := (*visited)[pos{next.x, next.y, next.dir, 0}]; ok && cost < next.cost {
				continue
			}

			(*visited)[pos{next.x, next.y, next.dir, 0}] = next.cost

			newPath := make([]pos, len(node.path))
			copy(newPath, node.path)
			newPath = append(newPath, next)

			queue = append(queue, Node{next, newPath})
		}
	}
	return
}
