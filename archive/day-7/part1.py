#!/usr/bin/python
import sys

def calculate(expected, arguments):
    a = arguments[0]
    b = arguments[1]
    tail = arguments[2:]

    return calc(expected, a*b, tail) + calc(expected, a+b, tail)

def calc(expected, result, tail):
    if len(tail) == 0:
        if result == expected:
            return 1
        else:
            return 0

    new_head = tail[0]
    new_tail = tail[1:]

    return calc(expected, result * new_head, new_tail) + calc(expected, result + new_head, new_tail)

     

sum = 0
for line in sys.stdin.readlines():
    line = line.rstrip()
    g = line.split(':')
    expected = int(g[0])
    arguments = list(map(lambda x: int(x), g[1].lstrip().split(' ')))
    print(expected)
    if calculate(expected, arguments) > 0:
        sum += expected

print(sum)

