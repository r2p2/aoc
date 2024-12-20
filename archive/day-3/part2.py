#!/usr/bin/python
import sys
import re

solution = 0
on = True
if __name__ == "__main__":
    input = sys.stdin.read()

    # input_new = re.sub(r'don\'t\(\).*?do\(\)', '', input)
    # print(input);
    # print(input_new);

    ops = re.findall(r'(mul\(([0-9]{1,3}),([0-9]{1,3})\)|don\'t\(\)|do\(\))', input)
    for group in ops:
        if group[0] == 'don\'t()':
            on = False
            continue;

        if group[0] == 'do()':
            on = True
            continue;

        if not on:
            continue

        solution += int(group[1])*int(group[2])
    
print(solution)


