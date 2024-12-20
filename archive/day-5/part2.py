#!/usr/bin/python
import sys

rules = []
pages = []
matches = []
nonmatches = []

in_rules = True
for line in sys.stdin.readlines():
    line = line.rstrip()
    if in_rules:
        if line == '':
            in_rules = False
            continue
        rules.append(line.split('|'))
        continue
    pages.append(line.split(','))

def is_matching(page):
    for rule in rules:
        before = rule[0]
        after = rule[1]
        if not (before in page and after in page):
            # rule does not apply
            continue

        idx_before = page.index(before)
        idx_after = page.index(after)

        if idx_before > idx_after:
            return False
    return True

def order(page):
    for i in range(len(page)):
        for r in rules:
            if r[0] != page[i]:
                continue
            if not r[1] in page[0:i]:
                continue

            # swap

            idx_before = page.index(r[1])
            idx_after = page.index(r[0])

            swp = page[idx_before]
            page[idx_before] = page[idx_after]
            page[idx_after] = swp
                

for page in pages:
    if not is_matching(page):
        nonmatches.append(page)

print("reorder")
for nm in nonmatches:
    print(nm)
    order(nm)
    order(nm)
    order(nm)
    order(nm)
    order(nm)
    order(nm)
    order(nm)
    order(nm)
    print(nm)


sum_nonmatch = 0
for m in nonmatches:
    middle = m[int((len(m)-1)/2)]
    sum_nonmatch += int(middle)
    print(middle)


print(rules)
print(pages)
print(nonmatches)
print(sum_nonmatch)

