import sys


def do_recurse(check_value, lines):
    bags = [line[1] for line in lines if check_value in line[0]]
    bags = bags.pop()
    mycount = 0
    for bag in bags.split(","):
        if bag != "no other":
            qty, mybag = bag.strip().split(" ", 1)
            qty = int(qty)
            val = do_recurse(mybag, lines)
            mycount += (val["count"] * qty) + qty
    return {"bags": [], "count": mycount}


def main():
    lines = []
    with open("input.txt") as input_file:
        for line in input_file:
            line = line.replace(".", "").replace("bags", "").replace("bag", "").strip()
            rule = line.split("contain")
            rule[0] = rule[0].strip()
            rule[1] = rule[1].strip()
            lines.append(rule)

    ans = do_recurse("shiny gold", lines)
    print(ans["count"])


if __name__ == "__main__":
    main(sys.argv[1:])
