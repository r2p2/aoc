#!/usr/bin/python
import fileinput

def is_safe(nx):
    deltas = []
    i = 0
    while i < len(nx)-1:
        deltas.append(int(nx[i+1]) - int(nx[i]))
        i += 1

    minimum = min(deltas)
    maximum = max(deltas)

    if 0 in deltas:
        return False

    if minimum < 0 and maximum > 0:
        return False

    if minimum < -3 or maximum > 3:
        return False

    return True

if __name__ == "__main__":
    safe = 0
    for line in fileinput.input():
        deltas = []
        nx = line.split()
        if is_safe(nx):
            safe += 1

    print(safe)


