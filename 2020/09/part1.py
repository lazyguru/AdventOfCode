import sys


def myfilter(pos, data):
    if pos < 25:
        return False
    last_25 = [int(lines[index]) for index in range(pos - 25, pos)]
    last_25.sort()
    while len(last_25) > 0:
        val = last_25.pop()
        if data - val in last_25:
            return False
    return True


def main():
    lines = []
    with open("input.txt") as input_file:
        lines = [line.strip() for line in input_file]

    good = [match for index, match in enumerate(lines) if myfilter(index, int(match))]
    print(good)


if __name__ == "__main__":
    main(sys.argv[1:])
