#!/usr/bin/env python3
import sys

def read_input(path):
    positions = []
    with open(path, 'r') as data:
        numbers = data.readline()
        for num in numbers.split(','):
            positions.append(int(num))
    return positions


def approximate_position(numbers):
    total, cost = 0, 0
    pairs = []
    for i in range(1, len(numbers)):
        m = (numbers[i] + numbers[i-1]) // 2
        pairs.append(m)

    for num in pairs:
        total += num
    return total//len(pairs), cost


   
def tweak_cost(numbers, positions, costs):
    stepmap = {}
    for i in range(len(positions)):
        for j in range(len(numbers)):
            steps = abs(numbers[j] - positions[i])
            try:
                costs[i] += stepmap[steps]
            except KeyError:
                f,g = 0,0
                for k in range(steps):
                    g += 1 + f
                    f += 1
                stepmap[steps] = g
                costs[i] += g 
    print(costs)
    print(stepmap)
    print(min(costs))

if __name__ == "__main__":
    numbers = read_input(sys.argv[1])
    print(numbers)
    position, cost = approximate_position(numbers)

    positions = [0] * len(numbers)
    costs = [0] * len(positions)

    for i in range(1, len(positions)):
        positions[i] = positions[i-1] + 1

    print(positions)
    tweak_cost(numbers, positions, costs)
