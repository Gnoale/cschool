#!/usr/bin/env python3
"""
TODO: handle negative memory access in every function (operation)
    op_diag 204,-1

"""


import sys


class Intcode:
    def __init__(self, data, start=0):
        self.run = data  # the intcode list
        self.start = start  # the first parameter
        self.pointer = 0
        self.opcode = None
        self.para1 = None
        self.para2 = None
        self.mode1 = None
        self.mode2 = None
        self.mode3 = None
        self.offset = 0

    def read_instructions(self):
        """
        Parameters in :
        Position mode (0) : the parameter value indicates the address from where to read the value
        Immediate mode (1): the parameter value is directly interpreted
        Relative mode (3) : the parameter value indicates the address relatively from the current base

        Relative base offset instruction change the relative base : opcode 9

        Mode 1, 2... give the parameters mode (position or immediate).
        Opcode is stored in the 2 rigth foremost digits, % 100 gives the remaining
        Mode 1 is stored in the 3th digit (hundreds)
        Mode 2 is stored in the 4th digit (thousand)

        Parameters are read from the next positions depending on the ops
        """
        opcodes = {
            1: "add",
            2: "mult",
            3: "store",
            4: "diag",
            5: "jumpt",
            6: "jumpf",
            7: "less",
            8: "equal",
            9: "offset",
            99: "exit",
        }

        self.mode1 = self.run[self.pointer] % 1000 // 100
        self.mode2 = self.run[self.pointer] % 10000 // 1000
        self.mode3 = self.run[self.pointer] % 100000 // 10000

        try:
            code = self.run[self.pointer] % 100
            self.opcode = opcodes[code]
        except KeyError:
            self.opcode = "error"
            print(f"Invalid opcode {code} : parameters mode = {self.mode1}, {self.mode2}, {self.mode3}")
            self.pointer += 2
            return

        # parameters are not interpreted for all instructions
        if self.opcode not in ["exit"]:
            self.para1 = self.get_parameters(self.mode1, 1)
            self.para2 = self.get_parameters(self.mode2, 2)
            self.para3 = self.get_parameters(self.mode3, 3)

        # print(self.run, self.pointer, self.run[self.para1], self.run[self.para2], self.run[self.para3])

    def get_parameters(self, mode, index):
        """Return the indice where to lookup the value based on the mode"""
        # positional
        if mode == 0:
            indice = self.memory_access(self.run[self.pointer + index])
        # immediate
        if mode == 1:
            indice = self.pointer + index
        # positional with offset
        if mode == 2:
            indice = self.memory_access(self.run[self.pointer + index] + self.offset)
        return indice

    def memory_access(self, index):
        """
        Sanity check on the index value
            If memory address > len(self.run) : return the modulus
            If memory address is < 0 : invalid, use the absolute value
        """
        try:
            if index < 0:
                index = abs(index)
            self.run[index]
            return index
        except IndexError:
            return index % len(self.run)

    def set_offset(self):
        """adjust the offset by the value of its only parameter"""
        self.offset += self.run[self.para1]
        self.pointer += 2

    def op_store(self):
        """store the "start" value at the address of the immediate and only parameter"""
        self.run[self.para1] = self.start
        self.pointer += 2

    def op_add(self):
        self.run[self.para3] = self.run[self.para1] + self.run[self.para2]
        self.pointer += 4

    def op_mult(self):
        self.run[self.para3] = self.run[self.para1] * self.run[self.para2]
        self.pointer += 4

    def op_diag(self):
        print(f"Diagnostic Code = {self.run[self.para1]}")
        self.pointer += 2

    def op_jumpf(self):
        if self.run[self.para1] == 0:
            self.pointer = self.run[self.para2]
        else:
            self.pointer += 3

    def op_jumpt(self):
        if self.run[self.para1] != 0:
            self.pointer = self.run[self.para2]
        else:
            self.pointer += 3

    def op_less(self):
        if self.run[self.para1] < self.run[self.para2]:
            self.run[self.para3] = 1
        else:
            self.run[self.para3] = 0
        self.pointer += 4

    def op_equal(self):
        if self.run[self.para1] == self.run[self.para2]:
            self.run[self.para3] = 1
        else:
            self.run[self.para3] = 0
        self.pointer += 4

    # run the intcode
    def op_run(self):
        while self.pointer < len(self.run):
            self.read_instructions()
            if self.opcode == "add":
                self.op_add()
            elif self.opcode == "mult":
                self.op_mult()
            elif self.opcode == "store":
                self.op_store()
            elif self.opcode == "diag":
                self.op_diag()
            elif self.opcode == "jumpt":
                self.op_jumpt()
            elif self.opcode == "jumpf":
                self.op_jumpf()
            elif self.opcode == "less":
                self.op_less()
            elif self.opcode == "equal":
                self.op_equal()
            elif self.opcode == "exit":
                return
            elif self.opcode == "offset":
                self.set_offset()
            elif self.opcode == "error":
                continue


if __name__ == '__main__':
    intcode = sys.argv[1]
    data = []
    with open(intcode, 'r') as fint:
        for line in fint:
            for code in line.split(','):
                data.append(int(code))

    machine = Intcode(data, 2)
    machine.op_run()
