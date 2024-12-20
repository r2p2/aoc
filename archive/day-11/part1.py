#!/usr/bin/python
import sys

line = sys.stdin.readlines()[0].rstrip()
print(line)

sgroups = line.split(' ')
groups = []
for e in sgroups:
    groups.append(int(e))

for i in range(25):
    print(i/75*100)
    new_group = []
    for e in groups:
        if e == 0:
            new_group.append(1)
        elif len(str(e)) % 2 == 0:
            se = str(e)
            new_group.append(int(se[0:int(len(se)/2)]))
            new_group.append(int(se[int(len(se)/2):]))
        else:
            new_group.append(e*2024)
    groups = new_group

print(len(groups))
