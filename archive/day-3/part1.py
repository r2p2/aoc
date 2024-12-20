#!/usr/bin/python
import sys
import re

solution = 0
if __name__ == "__main__":
    input = sys.stdin.read()
    ops = re.findall(r'mul\(([0-9]{1,3}),([0-9]{1,3})\)', input)
    for group in ops:
        solution += int(group[0])*int(group[1])
    
print(solution)


