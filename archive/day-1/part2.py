#!/usr/bin/python
import fileinput
from collections import defaultdict
lhs = []
rhs = defaultdict(int)
for line in fileinput.input():
    s = line.split()
    lhs.append(int(s[0]))
    rhs[int(s[1])] += 1

sum = 0
for n in lhs:
    sum += n * rhs[n]

print(sum)
    

