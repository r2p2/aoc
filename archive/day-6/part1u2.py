#!/usr/bin/python
import sys

def process(area, d, px, py):
    visited = set()
    while True:
        if (px, py, d) in visited:
            return (visited, True)
        visited.add((px,py,d))
        pxn = px
        pyn = py
        match d:
            case '^':
                pyn -= 1
            case 'v':
                pyn += 1
            case '<':
                pxn -= 1
            case '>':
                pxn += 1

        out_of_bounds = pxn < 0 or pyn < 0 or pxn >= width or pyn >= height
        if out_of_bounds:
            return (visited, False)

        if area[pyn][pxn] == '#':
            match d:
                case '^':
                    d = '>'
                case 'v':
                    d = '<'
                case '<':
                    d = '^'
                case '>':
                    d = 'v'
            continue
        px = pxn
        py = pyn

area_origin = []
width = 0
height = 0
d_orig = ' '
px_orig = 0
py_orig = 0
hits = 0

for line in sys.stdin.readlines():
    line = line.rstrip()
    width=len(line)
    height+=1
    area_origin.append(line)
    for x in range(len(line)):
        if line[x] in "^v<>":
            d_orig = line[x]
            px_orig = x
            py_orig = height-1

(visited, cycle) = process(area_origin, d_orig, px_orig, py_orig)

unique_visited = set()
for (x, y, d) in visited:
    if not (x,y) in unique_visited:
        unique_visited.add((x, y))

print("part 1:", len(unique_visited))

for n in range(1, len(unique_visited)):
        (x,y) = list(unique_visited)[n]
        area = area_origin.copy()
        area[y] = area[y][0:x] + '#' + area[y][x+1:]
        (visited, cycle) = process(area, d_orig, px_orig, py_orig)
        if cycle:
            hits += 1

print("part 2:", hits)


