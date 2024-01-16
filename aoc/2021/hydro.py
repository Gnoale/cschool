#!/usr/bin/env python3
import sys


def read_input(path):
    lines = []
    with open(path, 'r') as data:
        for line in data:
            p = line.split(' -> ')
            p0 = p[0].split(',')
            p1 = p[1].strip('\n').split(',')
            p0.extend(p1)
            lines.append(p0)
    return lines


def draw_points(lines):
    m = [[0]*1000 for i in range(1000)]
    for line in lines:
        x0 = int(line[0])
        x1 = int(line[2])
        y0 = int(line[1])
        y1 = int(line[3])
        # is horizontal
        if x0 == x1:
            j = y0
            if y0 < y1:
                while j <= y1:
                    m[x0][j] += 1
                    j += 1
            if y0 > y1:
                while j >= y1:
                    m[x0][j] += 1
                    j -= 1
        # vertical
        elif y0 == y1:
            i = x0
            if x0 < x1:
                while i <= x1:
                    m[i][y0] += 1
                    i += 1
            if x0 > x1:
                while i >= x1:
                    m[i][y0] += 1
                    i -= 1
        # diagonals
        elif x0 > x1 and y0 > y1:
            i = x0
            j = y0
            while i >= x1 and j >= y1:
                m[i][j] += 1
                i -= 1
                j -= 1
            
        elif x0 > x1 and y0 < y1:
            i = x0
            j = y0
            while i >= x1 and j <= y1:
                m[i][j] += 1
                i -= 1
                j += 1

        elif x0 < x1 and y0 < y1:
            i = x0
            j = y0
            while i <= x1 and j <= y1:
                m[i][j] += 1
                i += 1
                j += 1

        elif x0 < x1 and y0 > y1:
            i = x0
            j = y0
            while i <= x1 and j >= y1:
                m[i][j] += 1
                i += 1
                j -= 1
    return m


def count_inter(pmap):
    inter = 0
    for line in pmap:
        for point in line:
            if int(point) > 1:
                inter += 1
    print(inter)


if __name__ == "__main__":
    lines = read_input(sys.argv[1])
    # print(lines)
    m = draw_points(lines)
    count_inter(m)

