#!/usr/bin/python
import sys

area = []
width = 0
height = 0
d = ' '
px = 0
py = 0
visited=set()


for line in sys.stdin.readlines():
    line = line.rstrip()
    width=len(line)
    height+=1
    area.append(line)
    for x in range(len(line)):
        if line[x] in "^v<>":
            d = line[x]
            px = x
            py = height-1

print(area)
print(width)
print(height)
print(d)
print(px, py)

while True:
    print(" ", px, py, d)
    visited.add((px,py))
    pxn = px
    pyn = py
    if d == '^':
        pyn -= 1
    if d == 'v':
        pyn += 1
    if d == '<':
        pxn -= 1
    if d == '>':
        pxn += 1

    if pxn < 0 or pyn < 0 or pxn >= width or pyn >= height:
        break

    if area[pyn][pxn] == '#':
        if d == '^':
            d = '>'
            continue
        if d == 'v':
            d = '<'
            continue
        if d == '<':
            d = '^'
            continue
        if d == '>':
            d = 'v'
            continue
    px = pxn
    py = pyn

print(len(visited))


