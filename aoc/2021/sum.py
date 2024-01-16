#!/usr/bin/env python3
import sys


def read_input(path):
    data = []
    with open(path, 'r') as data_in:
        for num in data_in:
            data.append(int(num))
    return data


def count_depth(data):
    inc = 0
    for i in range(1, len(data)):
        if data[i] > data[i-1]:
            inc += 1
    return inc


def count_sliding_depth(data):
    inc = 0

    for i in range(2, len(data)):
        A = 0
        B = 0
        for j in range(i-2, i+1):
            A += data[j]

        try:
            for k in range(i-1, i+2):
                B += data[k]
            if B > A:
                inc += 1
        except IndexError:
            print('lol')

    return inc


if __name__ == "__main__":
    data = read_input(sys.argv[1])

    print(count_depth(data))
    print(count_sliding_depth(data))
