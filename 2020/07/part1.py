import sys


def do_recurse(check_value, lines, poss_bags):
    bags = [line[0] for line in lines if check_value in line[1]]
    poss_bags.update(set(bags))
    return [do_recurse(bag, lines, poss_bags) for bag in bags]


def main():
    lines = []
    with open("input.txt") as input_file:
        for line in input_file:
            line = line.replace(".", "").replace("bags", "").replace("bag", "").strip()
            rule = line.split("contain")
            rule[0] = rule[0].strip()
            rule[1] = rule[1].strip()
            lines.append(rule)

    poss_bags = set()
    do_recurse("shiny gold", lines, poss_bags)
    print(len(poss_bags))


if __name__ == "__main__":
    main(sys.argv[1:])
