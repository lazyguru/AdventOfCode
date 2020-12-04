import re


def is_valid_number_of_fields(passport):
    count = [1 for field in fields if field["id"] in passport]
    return len(count) == len(fields)


def is_valid_data(passport):
    count = [1 for field in fields if check_rule(field, passport)]
    return len(count) == len(fields)


def check_rule(field, data):
    rid = field["id"]
    rule = field["rules"]
    return rule.get("type", default)(rid, rule, data)


def numeric(rid, rule, data):
    if len(data[rid]) != rule["len"]:
        return False
    if int(data[rid]) < rule["min"]:
        return False
    if int(data[rid]) > rule["max"]:
        return False
    return True


def suffix(rid, rule, data):
    suf = data[rid][-2:]
    if suf not in rule["suffix"]:
        return False
    if int(data[rid][:-2]) < rule[f"min_{suf}"]:
        return False
    if int(data[rid][:-2]) > rule[f"max_{suf}"]:
        return False
    return True


def regex(rid, rule, data):
    return re.compile(rule["regex"]).match(data[rid])


def oneof(rid, rule, data):
    if data[rid] not in rule["oneof"]:
        return False
    return True


def default(rid, rule, data):
    return False


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
    good = [
        1
        for passport in passports
        if is_valid_number_of_fields(passport) and is_valid_data(passport)
    ]
    print(len(good))


fields = [
    {  # (Birth Year)
        "id": "byr",
        "rules": {"type": numeric, "len": 4, "min": 1920, "max": 2002},
    },
    {  # (Issue Year)
        "id": "iyr",
        "rules": {"type": numeric, "len": 4, "min": 2010, "max": 2020},
    },
    {  # (Expiration Year)
        "id": "eyr",
        "rules": {"type": numeric, "len": 4, "min": 2020, "max": 2030},
    },
    {  # (Height)
        "id": "hgt",
        "rules": {
            "type": suffix,
            "suffix": ["in", "cm"],
            "min_in": 59,
            "max_in": 76,
            "min_cm": 150,
            "max_cm": 193,
        },
    },
    {  # (Hair Color)
        "id": "hcl",
        "rules": {
            "type": regex,
            "regex": r"#[0-9a-z]{6}"
        }
    },
    {  # (Eye Color)
        "id": "ecl",
        "rules": {
            "type": oneof,
            "oneof": ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"],
        },
    },
    {  # (Passport ID)
        "id": "pid",
        "rules": {
            "type": numeric,
            "len": 9,
            "padding": True,
            "min": 1,
            "max": 999999999,
        },
    },
]

main()
