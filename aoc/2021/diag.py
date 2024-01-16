#!/usr/bin/env python3
import sys


def read_input(path):
    data = []
    with open(path, 'r') as data_in:
        for num in data_in:
            data.append(num.strip('\n'))
    return data


def most_common(data):
    """
    For each column (digit), record the most common occurence of 0 or 1
    return a mapping of each column and digit count 
        mapping[0] = [5,2] => the most occurence in first digit is 0
    """
    mapping = [0]*len(data[0])
    for i in range(len(mapping)):
        mapping[i] = [0, 0]
    for binary in data:
        for i in range(len(binary)):
            if int(binary[i]) == 0:
                mapping[i][0] += 1
            else:
                mapping[i][1] += 1
    return mapping


def get_rate(kind, bin_data):
    """
    "compute" a final binary number based on the most or less frequent occurence
    depending on the operation kind
    """
    final = ""
    if kind == "gamma":
        for mapping in bin_data:
            if mapping[0] > mapping[1]:
                final = final + "0"
            else:
                final = final + "1"
    elif kind == "epsilon":
        for mapping in bin_data:
            if mapping[0] < mapping[1]:
                final = final + "0"
            else:
                final = final + "1"
    return final


def get_numbers(kind, bin_data, data, index):
    """
    For each most or less common bit, depending of the kind, 
    keep only the numbers that matches the bit in the current index.
    """
    numbers = []
    if bin_data[index][0] > bin_data[index][1]:
        most_common = 0
        least_common = 1
    else:
        most_common = 1
        least_common = 0
    if kind == "o2":
        for binary in data:
            if int(binary[index]) == most_common:
                numbers.append(binary)
    elif kind == "co2":
        for binary in data:
            if int(binary[index]) == least_common:
                numbers.append(binary)
    return numbers


if __name__ == "__main__":
    data = read_input(sys.argv[1])

    bin_data = most_common(data)
    print(int(get_rate("gamma", bin_data), 2) *
          int(get_rate("epsilon", bin_data), 2))

    print(bin_data)

    o2 = data.copy()
    co2 = data.copy()
    o2_map = bin_data.copy()
    co2_map = bin_data.copy()

    for index in range(len(bin_data)):
        print(index)
        if len(o2) > 1:
            o2 = get_numbers("o2", o2_map, o2, index)
            o2_map = most_common(o2)
        if len(co2) > 1:
            co2 = get_numbers("co2", co2_map, co2, index)
            co2_map = most_common(co2)
        print(o2, co2)
        print(o2_map, co2_map)

    print(int(o2[0], 2)*int(co2[0], 2))
