#!/usr/bin/env python3
import sys

def read_input(path):
    data = []
    with open(path, 'r') as data_in:
        for cmd in data_in:
            data.append(cmd.strip('\n'))
    return data


def read_position(data):

    position = [0,0]
    total = 0

    for moves in data:
        move = moves.split()
        direction = move[0]
        step = int(move[1])

        if direction == "forward":
            position[0] += step
        elif direction == "down":
            position[1] += step
        elif direction == "up":
            position[1] -= step
    
    return position, position[0] * position[1]
        

def read_position_aim(data):

    position = [0,0]
    total = 0
    aim = 0

    for moves in data:
        move = moves.split()
        direction = move[0]
        step = int(move[1])

        if direction == "forward":
            position[0] += step
            if aim != 0:
                position[1] += aim * step
        elif direction == "down":
            aim += step
        elif direction == "up":
            aim -= step
    
    return position, position[0] * position[1]
        

if __name__ == "__main__":
    data = read_input(sys.argv[1])

    print(read_position(data))
    print(read_position_aim(data))
