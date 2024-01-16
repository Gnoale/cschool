#!/usr/bin/env python3
import sys, re

def read_input(path):
    # light red bags contain 1 bright white bag, 2 muted yellow bags.
    rules = {}
    with open(path, 'r') as data:
        for line in data:
            tmp = line.split(' bags contain ')
            t = []
            for b in tmp[1].split(','):
                b = re.sub(' bag.*', '', b).strip('\n')
                t.append(b.lstrip(' '))
            rules[tmp[0]] = t 
    return rules

def count_colors(rules, color, visited):

    if len(visited) == len(rules):
        return

    for parent, colors in rules.items():
        try:
            visited[parent]
        except KeyError:
            for bag in colors:
                if color == bag[2:]:
                    visited[parent] = 1
                    count_bag(rules, parent, visited)
    

def count_bags(rules, color, visited, factor):
    
    if len(visited) == len(rules):
        return

    childs = rules[color]
    for child in childs:
        ccolor = child[2:]
        if ccolor == ' other':
            continue
        ccount = int(child[0])
        print(color, visited, factor)
        try:
            visited[ccolor] += ccount * factor
        except KeyError:
            visited[ccolor] = ccount * factor
        count_bags(rules, ccolor, visited, ccount*factor)
    


if __name__ == "__main__":
    rules = read_input(sys.argv[1])
    visited = {}
    #count_colors(rules, 'shiny gold', visited)
    print(len(visited))

    visited = {}
    count_bags(rules, 'shiny gold', visited, 1)
    print(visited)
    total = 0
    for k,v in visited.items():
        total += v
    print(total)
