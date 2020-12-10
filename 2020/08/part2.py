import sys


def reset_seen(data):
    data["seen"] = False
    return data


def tryloop(ch_pos, ch_op, lines):
    ans = 0
    pos = 0
    while pos < len(lines):
        if lines[pos]["seen"]:
            return False
        lines[pos]["seen"] = True
        amt = lines[pos]["amt"]
        oper = lines[pos]["op"]
        if pos == ch_pos:
            oper = ch_op
        if oper == "jmp":
            pos += amt
            continue
        if oper == "nop":
            pos += 1
            continue
        if oper == "acc":
            ans += amt
            pos += 1
            continue
    print(ans)
    return True


def create_data(data):
    oper, amt = data.replace("+", "").split(" ")
    return {"op": oper, "amt": int(amt), "seen": False}


def main():
    lines = []
    with open("input.txt") as input_file:
        lines = [create_data(line) for line in input_file]

    nops = [int(index) for index, match in enumerate(lines) if match["op"] == "nop"]
    jmps = [int(index) for index, match in enumerate(lines) if match["op"] == "jmp"]
    for nop in nops:
        for line in lines:
            reset_seen(line)
        if tryloop(nop, "jmp", lines):
            return

    for jmp in jmps:
        for line in lines:
            reset_seen(line)
        if tryloop(jmp, "nop", lines):
            return


if __name__ == "__main__":
    main(sys.argv[1:])
