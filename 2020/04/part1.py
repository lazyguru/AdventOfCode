def is_valid(passport):
    count = [1 for field in fields if field in passport]
    return len(count) == len(fields)


def main():
    passports = []
    with open("input.txt") as input_file:
        passport = {}
        for line in input_file:
            line = line.strip()
            if line == "":
                passports.append(passport)
                passport = {}
                continue
            flds = line.split(" ")
            data = [fld.split(":") for fld in flds]
            for item in data:
                key = item[0].strip()
                value = item[1].strip()
                passport[key] = value

    good = [True for passport in passports if is_valid(passport)]
    print(len(good))


fields = [
    "byr",  # (Birth Year)
    "iyr",  # (Issue Year)
    "eyr",  # (Expiration Year)
    "hgt",  # (Height)
    "hcl",  # (Hair Color)
    "ecl",  # (Eye Color)
    "pid",  # (Passport ID)
]

main()
