import math


def myfilter(data):
    cols = len(data["line"])
    # The first row does count
    if data["idx"] == 0:
        return False
    pos = data["idx"] * 3
    # If we move past the right boundary, wrap
    if pos >= cols:
        # we need to account for how many wraps have happened
        pos -= math.floor(pos / cols) * cols
    print(f"{data['idx']}:{pos}:{cols}")
    print(data["line"])
    spot = data["line"][pos]
    print(spot)
    return spot == "#"


lines = []
with open("input.txt") as f:
    i = 0
    for line in f:
        lines.append({"idx": i, "line": line.strip()})
        i += 1

good = [line["idx"] for line in lines if myfilter(line)]
print(len(good))
