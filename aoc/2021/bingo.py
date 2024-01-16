#!/usr/bin/env python3
import sys

def read_input(path):
    boards = []
    numbers = []

    with open(path, 'r') as data:
        numbers = [int(number) for number in data.readline().strip('\n').split(',')]
        unparsed_board = data.read().split()
    line = []
    board = []
    i = 1
    for number in unparsed_board:
        if i % 25 == 0:
            line.append(int(number))
            board.append(line)
            boards.append(board)
            line = []
            board = []
            i += 1
            continue
        if i % 5 == 0:
            line.append(int(number))
            board.append(line)
            line = []
        else:
            line.append(int(number))
        i += 1
    return (numbers, boards)


def run_board(boards, numbers, wins):
    bw = [False] * len(boards)
    for i in range(len(numbers)):
        for j in range(len(boards)):
            if not bw[j]:
                play_board(numbers[i], boards[j])
                win, total = board_win(boards[j])
                if win:
                    wins[i] = total*numbers[i]
                    print(wins[i], total,  numbers[i], boards[j])
                    bw[j] = True

def play_board(number, board):
    for line in board:
        for i in range(len(line)):
            if line[i] == number:
                line[i] = True

def board_win(board):
    total = 0
    column = [0]*5
    line = [0]*5
    for i in range(len(board)):
        for j in range(len(board[i])):
            if board[i][j] is True:
                line[i] += 1
                column[j] += 1
            else:
                total += board[i][j]
    for count in line:
        if count == 5:
            return True, total
    for count in column:
        if count == 5:
            return True, total

    return False, total

if __name__ == "__main__":
    numbers, boards = read_input(sys.argv[1])
    
    wins = [0] * len(boards)
    #for board in boards:
    run_board(boards, numbers, wins)

    print(wins)
