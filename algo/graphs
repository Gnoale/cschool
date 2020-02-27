#!/usr/bin/env python

from collections import deque, defaultdict

# dag
orbitals = [ 
    ('C','D'),
    ('B','C'),
    ('D','E'),
    ('B','G'),
    ('E','F'),
    ('G','H'),
    ('D','I'),
    ('I', 'SAN'),
    ('E','J'),
    ('J','K'),
    ('K', 'YOU'),
    ('COM','B'),
    ('K','L')
    ]

# Is this entirely running in O(1) ?
def bfs(node, graph):
    stack = deque()
    stack.append(node)
    visited = []
    level = {}
    visited.append(node)
    level[node] = 0
    while len(stack) > 0:
        # select the node to visit
        node = stack.pop()
        next_nodes = graph[node]
        for n in next_nodes:
            if n not in visited:
                # track the graph level
                level[n] = level[node] + 1
                stack.append(n)
                visited.append(n)
    return level


parent = {}
def dfs(graph):
    for node,childs in graph.items():
        if node not in parent:
            dfs_visit(node,childs,graph)

def dfs_visit(node,childs,graph):
    print(childs)
    for c in childs:
        if c not in parent:
            parent[c] = node
            dfs_visit(c,graph[c],graph)




if __name__ == '__main__':
    
    # set the graph datastructure
    graph = defaultdict(list)
    for k, v in orbitals:
        graph[k].append(v)

    # traverse the graph with bfs
    level = bfs('COM',graph)
    print(level,'\n')

    # traverse the graph with dfs
    dfs(graph)
    for k,v in parent.items():
        print("{} parent = {}".format(k,v))
