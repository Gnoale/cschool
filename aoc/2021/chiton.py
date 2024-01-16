#!/usr/bin/env python3
import sys
from collections import defaultdict


def read_input(path):
    maps = []
    with open(path, 'r') as data_in:
        for data in data_in:
            maps.append([int(d) for d in data.strip('\n')])
    return maps


def graph(maps):
    g = defaultdict(list)
    for i in range(len(maps)):
        for j in range(len(maps[i])):
            if j == len(maps[i]) - 1 and i == len(maps) - 1:
                return g
            elif j == len(maps[i]) - 1:
                g[(i, j)].append((i+1, j))
            elif i == len(maps) - 1:
                g[(i, j)].append((i, j+1))
            else:
                g[(i, j)].append((i+1, j))
                g[(i, j)].append((i, j+1))


# def traverse_graph(node, graph):
#    queue = []
#    visited = []
#    level = {}
#
#    queue.append(node)
#    visited.append(node)
#    level[node] = 0
#
#    paths = []
#    paths.append(node)
#
#    while len(queue) > 0:
##        print("Visited : ",visited)
#        node = queue.pop(0)
#        print(node)
#        next_nodes = graph[node]
#        for n in next_nodes:
#            paths.append(node)
#            level[n] = level[node] + 1
#            if n not in visited:
#                queue.append(n)
#                visited.append(n)
#
#    return level

def dfs(graph):
    for node,childs in graph.items():
        if node not in parent:
            dfs_visit(node,childs,graph)

def dfs_visit(node, childs, graph):
    for c in childs:
        print(c)
        print(parent)
        if c not in parent:
            parent[c] = node
            dfs_visit(c, graph[c], graph)

parent = {}



maps = read_input(sys.argv[1])
print(maps)

g = graph(maps)

#level = traverse_graph((0,0), g)
print("")
print(g)

print(len(g))
dfs(g)

#get_shortest(level, maps)
