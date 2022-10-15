import sys


def myfilter(data):
    return False


lines = []
with open("input.txt") as input_file:
    lines = [line for line in input_file]

good = [1 for line in lines if myfilter(line)]
print(len(good))
