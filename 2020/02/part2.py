import re


def myfilter(data):
    pos1,pos2,char,ignore,pwd,_ = data
    pos1 = int(pos1) - 1 # handle for 0 index
    pos2 = int(pos2) - 1 # handle for 0 index
    if pwd[pos1:pos1+1] == char and pwd[pos2:pos2+1] == char:
        return False
    return char in [pwd[pos1:pos1+1], pwd[pos2:pos2+1]]

lines = []
regex = re.compile(r'\W')
with open('input.txt') as f:
    for line in f:
        lines.append(regex.split(line))

good = [ line[4] for line in lines if myfilter(line) ]
print(len(good))
