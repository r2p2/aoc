#!/usr/bin/python
import fileinput
lhs = []
rhs = []
for line in fileinput.input():
    s = line.split()
    lhs.append(int(s[0]))
    rhs.append(int(s[1]))

lhs.sort()
rhs.sort()

sum = 0
for l, r in zip(lhs, rhs):
    sum += abs(l-r)
print(sum)
    

