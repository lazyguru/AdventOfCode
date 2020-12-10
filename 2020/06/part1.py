import sys


def main():
    groups = []
    with open("input.txt") as input_file:
        group = set()
        for line in input_file:
            line = line.strip()
            if line == "":
                groups.append(group)
                group = set()
                continue
            group = group.union(set(line))

    count = [len(group) for group in groups]
    print(sum(count))


if __name__ == "__main__":
    main(sys.argv[1:])
