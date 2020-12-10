import sys

ROWS = 127
COLS = 7


def parse_seat(seat_data):
    high = ROWS
    low = 0
    for pos in seat_data[0:6]:
        if pos == "F":
            high -= round((high - low) / 2)
        if pos == "B":
            low += round((high - low) / 2)
    my_row = high
    if seat_data[6] == "F":
        my_row = low
    high = COLS
    low = 0
    for pos in seat_data[7:9]:
        if pos == "L":
            high -= round((high - low) / 2)
        if pos == "R":
            low += round((high - low) / 2)
    my_seat = high
    if seat_data[9] == "L":
        my_seat = low
    return (my_row * 8) + my_seat


def main():
    lines = []
    with open("input.txt") as input_file:
        lines = [line.strip() for line in input_file]

    seat_ids = [parse_seat(seat) for seat in lines]
    seat_ids.sort()
    """
    This needs to be cleaned up more. Right now it prints a bunch of values
    Ideally, it should only print your ticket number
    """
    res = list(set(range(max(seat_ids) + 1)) - set(seat_ids))
    print(res)


if __name__ == "__main__":
    main(sys.argv[1:])
