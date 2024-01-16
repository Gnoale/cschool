#!/usr/bin/env python3
import sys, os
from multiprocessing import Process, Manager

def read_input(path):
    with open(path, 'r') as data:
        return [int(state) for state in data.readline().strip('\n').split(',')]


class Fish:

    def __init__(self, age):
        self.age = age

    def is_aging(self):
        if self.age == 0:
            self.age = 6
            return True
        self.age -= 1
        return False


def reboot_and_fork(fishmap, factor):
    # reboot
    fishmap[0] -= 1 * factor
    fishmap[6] += 1 * factor
    # fork
    fishmap[8] += 1 * factor


def age_fish(fishmap, age, factor):
    fishmap[age] -= 1 * factor
    fishmap[age-1] += 1 * factor

def simulate(fishmap, days):
    if days == 0:
        return
    # instead of tracking all fishes
    # create one per age and use the count of fish to track ages
    fmap = fishmap.copy()
    for age, count in fmap.items():
        if count > 0:
            if age == 0:
                reboot_and_fork(fishmap, count)
            else:
                age_fish(fishmap, age, count)
    days -= 1
    simulate(fishmap, days)


if __name__ == "__main__":
    initial = read_input(sys.argv[1])

    fishmap = {i:0 for i in range(9)}
    
    for age in initial:
        fishmap[age] += 1

    print(fishmap)
    simulate(fishmap, 256)

    total = 0
    for k,v in fishmap.items():
        total += v
    print(total)
   

