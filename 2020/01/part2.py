import sys


nums = []

with open('input.txt') as f:
    for line in f:
        nums.append(int(line.strip()))

nums.sort()

while len(nums) > 0:
    x = nums.pop(1)
    res = [ (2020 - (x + y)) * x * y for y in nums if 2020 - (x + y) in nums ]
    if res:
        print(res.pop())
        sys.exit()
