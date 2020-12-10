def create_data(data):
    oper, val = data.replace("+", "").split(" ")
    return {"op": oper, "amt": int(val), "seen": False}


def main():
    lines = []
    with open("input.txt") as input_file:
        lines = [create_data(line) for line in input_file]

    ans = 0
    pos = 0
    while True:
        if lines[pos]["seen"]:
            break
        lines[pos]["seen"] = True
        amt = lines[pos]["amt"]
        if lines[pos]["op"] == "jmp":
            pos += amt
            continue
        if lines[pos]["op"] == "nop":
            pos += 1
            continue
        if lines[pos]["op"] == "acc":
            ans += amt
            pos += 1
            continue
    print(ans)


main()
