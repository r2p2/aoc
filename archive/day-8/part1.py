#!/usr/bin/python
import sys

antennas = {}
width = 0
height = 0
antinodes = []

for line in sys.stdin.readlines():
    height += 1
    line = line.rstrip()
    width = len(line)
    for x in range(len(line)):
        if line[x] == '.':
            continue

        if not line[x] in antennas:
            antennas[line[x]] = []

        antennas[line[x]].append((x, height - 1))

# calc atifoos
for freq in antennas.keys():
    print(freq)

    for i in range(len(antennas[freq])):
        for j in range(len(antennas[freq])):
            if i == j:
                continue
            a = antennas[freq][i]
            b = antennas[freq][j]
            d = (abs(b[0] - a[0]), abs(b[1] - a[1]))

            nx = a[0]
            ny = a[1]
            if a[0] < b[0]:
                nx -= d[0]
            if a[0] > b[0]:
                nx += d[0]
            if a[1] < b[1]:
               ny -= d[1]
            if a[1] > b[1]:
                ny += d[1]

            n = (nx, ny)

            print(a, "->", b, "d:", d, "n:", n)
            if nx < 0 or ny < 0 or nx >= width or ny >= height:
                continue

            if not n in antinodes:
                antinodes.append(n)
    


print(antinodes)
print(len(antinodes))
