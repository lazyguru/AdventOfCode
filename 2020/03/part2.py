import sys
import math


def myfilter(data, slope):
    cols = len(data["line"])
    # The first row does count
    if data["idx"] == 0:
        return False
    pos = data["idx"] * slope
    # If we move past the right boundary, wrap
    if pos >= cols:
        # we need to account for how many wraps have happened
        pos -= math.floor(pos / cols) * cols
    print(f"{data['idx']}:{pos}:{cols}")
    print(data["line"])
    spot = data["line"][pos]
    print(spot)
    return spot == "#"


def main():
    lines = []
    with open("input.txt") as f:
        i = 0
        for line in f:
            lines.append({"idx": i, "line": line.strip()})
            i += 1

    slope31 = [line["idx"] for line in lines if myfilter(line, 3)]
    slope11 = [line["idx"] for line in lines if myfilter(line, 1)]
    slope51 = [line["idx"] for line in lines if myfilter(line, 5)]
    slope71 = [line["idx"] for line in lines if myfilter(line, 7)]
    slope12 = [
        line["idx"] for line in lines if line["idx"] % 2 == 0 and myfilter(line, 1)
    ]

    print(len(slope31) * len(slope11) * len(slope51) * len(slope71) * len(slope12))


if __name__ == "__main__":
    main(sys.argv[1:])
