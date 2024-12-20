#!/usr/bin/python
import sys
lines = sys.stdin.readlines()
matches = 0

for y in range(1, len(lines)-1): 
    for x in range(1, len(lines[0])-1):
        if lines[y][x] != 'A':
            continue

        try:
            a = {'M': 0, 'S': 0}
            a[lines[y-1][x-1]] += 1
            a[lines[y+1][x+1]] += 1

            b = {'M': 0, 'S': 0}
            b[lines[y+1][x-1]] += 1
            b[lines[y-1][x+1]] += 1

            if a['M'] == 1 and a['S'] == 1 and b['M'] == 1 and b['S'] == 1:
                matches += 1
        except KeyError:
            pass

print(matches)
