import sys
import re


def myfilter(data):
    low, high, pattern, ignore, pwd, _ = data
    match = re.findall(pattern, pwd)
    if match is None:
        return False
    return len(match) <= int(high) and len(match) >= int(low)


def main():
    lines = []
    regex = re.compile(r"\W")
    with open("input.txt") as f:
        for line in f:
            lines.append(regex.split(line))

    good = [line[4] for line in lines if myfilter(line)]
    print(len(good))


if __name__ == "__main__":
    main(sys.argv[1:])
