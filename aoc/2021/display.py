#!/usr/bin/env python3
import sys
import re
from itertools import permutations

def read_input(path):
    patterns = []
    with open(path, 'r') as data_in:
        for data in data_in:
            patterns.append([d.strip(' \n') for d in data.split('|')])
    return patterns

def count_unique_patterns(numbers, patterns):
    total = 0
    for pattern in patterns:
        # map unique numbers permutations
        m = {}
        #   0
        # 3   1
        #   2
        # 4   5
        #   6
        for digit in pattern[0].split():
            # ['fcgedb', 'cgb', 'dgebacf', 'gc']
            # ['cg', 'cg', 'fdcagb', 'cbg']
            # get pattern from number representations that are uniques
            num = numbers[len(digit)]
            # we got all permutations for each uniques number
            if len(num) == 1:
                m[num[0]] = list(permutations([d for d in digit]))

        # find others set
        s15 = set(m[1][0])
        s1235 = set(m[4][0])
        s015 = set(m[7][0])
        s0123456 = set(m[8][0])

        s0 = s015 - s15
        s2346 = s0123456 - s015
        s046 = s0123456 - s1235
        s46 = s046 - s0
        s23 = s2346 - s46
        # generates best permutations for other numbers
        # 2 = 01246 = s0 + s46 = 3
        # 3 = 01256 = s0 + s15 = 3
        # 5 = 02356 = s0 + s23 = 3
        # 6 = 023456 = s0 + s2346 = 5
        # 9 = 012356 = s015 + s23 = 5
        # 0 = 013456 = s015 + s46 = 5
        p2 = s0.union(s46)
        p3 = s015
        p5 = s0.union(s23)
        p6 = s0.union(s2346)
        p9 = s015.union(s23)
        p0 = s015.union(s46)
        # now find the numbers !
        num = ""
        for digit in pattern[1].split():
            n = len(digit)
            ln = numbers[n]
            # easy case
            if len(ln) == 1:
                num += str(ln[0])
            # find best match depending on digit numbers
            else:
                if n == 5:
                    num += find_best(digit, 3, {"2": p2, "3": p3, "5": p5})
                if n == 6:
                    num += find_best(digit, 5, {"0": p0, "6": p6, "9": p9})

        total += int(num)
    print(total)


def find_best(digit, best, permutations):
    # find best match among the list or permutations
    num = ""
    for number, permutation in permutations.items():
        cpt = 0
        for char in permutation:
            if char in digit:
                cpt += 1
            if cpt == best:
                num = number
                break


    if num == "":
        print("fail",digit, best)
        print(permutations)
    return num

if __name__ == "__main__":
    patterns = read_input(sys.argv[1])
    # index numbers by their segments count
    numbers = {2:[1], 5:[2,3,5], 4:[4], 6:[0,6,9], 3:[7], 7:[8]}
    count_unique_patterns(numbers, patterns)

    # model our display
    dp = ["a","b","c","d","e","f","g"]
