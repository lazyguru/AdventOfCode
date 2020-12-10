import sys


def main():
    groups = []
    with open("input.txt") as input_file:
        group = {"count": 0}
        for line in input_file:
            line = line.strip()
            if line == "":
                g = [
                    {question: group[question]}
                    for question in group
                    if group[question] == group["count"] and question != "count"
                ]
                groups.append(g)
                group = {"count": 0}
                continue
            group["count"] += 1
            for char in line:
                if char in group:
                    group[char] += 1
                else:
                    group[char] = 1
    count = [len(group) for group in groups]
    print(sum(count))


if __name__ == "__main__":
    main(sys.argv[1:])
