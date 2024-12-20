#!/usr/bin/python
import sys

line = sys.stdin.readlines()[0].rstrip()
print(line)
solution = {}

sgroups = line.split(' ')
groups = []
for e in sgroups:
    groups.append(int(e))

for i in range(75):
    print(i/75*100)
    new_group = []

    for e in groups:

        new_es = []
        if e in sgroups:
            new_es = sgroups[e]
        else:
            if e == 0:
                new_es.append(1)
            elif len(str(e)) % 2 == 0:
                se = str(e)
                new_es.append(int(se[0:int(len(se)/2)]))
                new_es.append(int(se[int(len(se)/2):]))
            else:
                new_es.append(e*2024)
        solution[e] = new_es
        for n in new_es:
            new_group.append(n)
    groups = new_group

print(len(groups))
