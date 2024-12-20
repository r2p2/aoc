#!/usr/bin/python
import sys

rules = []
pages = []
matches = []

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

for page in pages:
    fail = False
    
    for rule in rules:
        before = rule[0]
        after = rule[1]
        if not (before in page and after in page):
            # rule does not apply
            continue

        idx_before = page.index(before)
        idx_after = page.index(after)

        if idx_before > idx_after:
            fail = True

    if fail:
        continue

    matches.append(page)

sum = 0
for m in matches:
    middle = m[int((len(m)-1)/2)]
    sum += int(middle)
    print(middle)


print(rules)
print(pages)
print(matches)
print(sum)

