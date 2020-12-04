nums = []

with open("input.txt") as f:
    for line in f:
        nums.append(int(line.strip()))

res = [(2020 - x) * x for x in nums if 2020 - x in nums]
print(res.pop())
