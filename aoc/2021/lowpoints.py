#!/usr/bin/env python3
import sys
from dataclasses import dataclass

MAPS = []

def read_input(path):
    with open(path, 'r') as data_in:
        for data in data_in:
            MAPS.append([int(d) for d in data.strip('\n')])


def find_lower():
    lowers = []
    for i in range(len(MAPS)):
        for j in range(len(MAPS[i])):
            adjacency = get_adjacencies(i, j)
            # check if current point is lower than adjacencies
            low = True
            for adj in adjacency:
                if MAPS[i][j] >= adj:
                    low = False
            if low:
                lowers.append(MAPS[i][j]+1)
    total = 0
    for lower in lowers:
        total += lower
    print(total)
    return lowers


def get_adjacencies(i,j):
    adjacency = []
    # corners
    if i == 0 and j == 0:
        adjacency.append(MAPS[i][j+1])
        adjacency.append(MAPS[i+1][j])
    elif i == 0 and j == len(MAPS[i]) - 1:
        adjacency.append(MAPS[i][j-1])
        adjacency.append(MAPS[i+1][j])
    elif i == len(MAPS) - 1 and j == 0:
        adjacency.append(MAPS[i][j+1])
        adjacency.append(MAPS[i-1][j])
    elif i == len(MAPS) - 1 and j == len(MAPS[i]) - 1:
        adjacency.append(MAPS[i][j-1])
        adjacency.append(MAPS[i-1][j])
    # edges
    elif j == 0:
        adjacency.append(MAPS[i-1][j])
        adjacency.append(MAPS[i][j+1])
        adjacency.append(MAPS[i+1][j])
    elif j == len(MAPS[i]) - 1:
        adjacency.append(MAPS[i-1][j])
        adjacency.append(MAPS[i][j-1])
        adjacency.append(MAPS[i+1][j])
    elif i == 0:
        adjacency.append(MAPS[i][j-1])
        adjacency.append(MAPS[i][j+1])
        adjacency.append(MAPS[i+1][j])
    elif i == len(MAPS) - 1:
        adjacency.append(MAPS[i][j-1])
        adjacency.append(MAPS[i][j+1])
        adjacency.append(MAPS[i-1][j])
    # rest of the map
    else:
        adjacency.append(MAPS[i][j-1])
        adjacency.append(MAPS[i][j+1])
        adjacency.append(MAPS[i-1][j])
        adjacency.append(MAPS[i+1][j])
    return adjacency


def get_adjacency_points(point):
    points = []
    i = point.i
    j = point.j
    # corners
    if i == 0 and j == 0:
        points.append(Point(i, j+1))
        points.append(Point(i+1, j))
    elif i == 0 and j == len(MAPS[i]) - 1:
        points.append(Point(i, j-1))
        points.append(Point(i+1, j))
    elif i == len(MAPS) - 1 and j == 0:
        points.append(Point(i, j+1))
        points.append(Point(i-1, j))
    elif i == len(MAPS) - 1 and j == len(MAPS[i]) - 1:
        points.append(Point(i, j-1))
        points.append(Point(i-1, j))
    # edges
    elif j == 0:
        points.append(Point(i-1, j))
        points.append(Point(i, j+1))
        points.append(Point(i+1, j))
    elif j == len(MAPS[i]) - 1:
        points.append(Point(i-1, j))
        points.append(Point(i, j-1))
        points.append(Point(i+1, j))
    elif i == 0:
        points.append(Point(i, j-1))
        points.append(Point(i, j+1))
        points.append(Point(i+1, j))
    elif i == len(MAPS) - 1:
        points.append(Point(i, j-1))
        points.append(Point(i, j+1))
        points.append(Point(i-1, j))
    # rest of the map
    else:
        points.append(Point(i, j-1))
        points.append(Point(i, j+1))
        points.append(Point(i-1, j))
        points.append(Point(i+1, j))
    return points


@dataclass
class Point:
    i: int
    j: int

    def is_bassin(self):
        return MAPS[self.i][self.j] < 9

    def value(self):
        return MAPS[self.i][self.j]

def find_bassins():
    bassins = []
    visited = []
    for i in range(len(MAPS)):
        for j in range(len(MAPS[i])):
            bassin = visit(i,j,visited)
            if len(bassin) > 1:
                bassins.append(bassin)

    L = []
    for bassin in bassins:
        L.append(len(bassin))

    L.sort()
    print(L.pop() * L.pop() * L.pop())

def visit(i,j, visited):
    bassin = []
    point = Point(i,j)
    if point in visited:
        return []
    if point.is_bassin():
        to_visit = [point]
        bassin.append(point)
        while len(to_visit) > 0:
            point = to_visit.pop()
            visited.append(point)
            adj_points = get_adjacency_points(point)
            for p in adj_points:
                if p.is_bassin() and p not in bassin:
                    bassin.append(p)
                    to_visit.append(p)

    return bassin



if __name__ == "__main__":
    read_input(sys.argv[1])
    find_lower()
    print("### bassins")
    find_bassins()
