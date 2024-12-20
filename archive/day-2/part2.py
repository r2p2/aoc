#!/usr/bin/python
import fileinput
from part1 import is_safe

def is_safe2(nx):
    if is_safe(nx):
        return True;
    i=0
    while i < len(nx):
        nx2 = nx.copy();
        nx2.pop(i)
        i += 1
        if is_safe(nx2):
            return True
    return False

safe = 0
for line in fileinput.input():
    deltas = []
    nx = line.split()
    if is_safe2(nx):
        safe += 1

print(safe)


