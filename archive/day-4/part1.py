#!/usr/bin/python
import fileinput
lines = []
matches = 0
for line in fileinput.input():
    lines.append(line);

for y in range(len(lines)): 
    for x in range(len(lines[0])):
       if lines[y][x] != 'X':
           continue

       try:
           if lines[y+1][x] == 'M' and lines[y+2][x] == 'A' and lines[y+3][x] == 'S':
               matches += 1
       except IndexError:
           pass

       try:
           if y > 2 and lines[y-1][x] == 'M' and lines[y-2][x] == 'A' and lines[y-3][x] == 'S':
               matches += 1
       except IndexError:
           pass

       try:
           if lines[y][x+1] == 'M' and lines[y][x+2] == 'A' and lines[y][x+3] == 'S':
               matches += 1
       except IndexError:
           pass

       try:
           if x > 2 and lines[y][x-1] == 'M' and lines[y][x-2] == 'A' and lines[y][x-3] == 'S':
               matches += 1
       except IndexError:
           pass

       try:
           if lines[y+1][x+1] == 'M' and lines[y+2][x+2] == 'A' and lines[y+3][x+3] == 'S':
               matches += 1
       except IndexError:
           pass

       try:
           if x > 2 and y > 2 and lines[y-1][x-1] == 'M' and lines[y-2][x-2] == 'A' and lines[y-3][x-3] == 'S':
               matches += 1
       except IndexError:
           pass

       try:
           if y > 2 and lines[y-1][x+1] == 'M' and lines[y-2][x+2] == 'A' and lines[y-3][x+3] == 'S':
               matches += 1
       except IndexError:
           pass

       try:
           if x > 2 and lines[y+1][x-1] == 'M' and lines[y+2][x-2] == 'A' and lines[y+3][x-3] == 'S':
               matches += 1
       except IndexError:
           pass

print(matches)
