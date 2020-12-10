import sys


def find_num(invalid_num, lines):
    pos = 0
    while pos < len(lines):
        ans = invalid_num
        my_pos = pos
        my_ans = []
        while ans >= 0:
            val = int(lines[my_pos])
            ans -= val
            my_ans.append(val)
            my_pos += 1
            if ans == 0:
                return my_ans
        pos += 1


def main():
    lines = []
    with open("input.txt") as input_file:
        lines = [line.strip() for line in input_file]

    lines.remove("22477624")
    good = find_num(22477624, lines)
    good.sort()
    answer = good[0] + good[len(good) - 1]
    print(answer)


if __name__ == "__main__":
    main(sys.argv[1:])
