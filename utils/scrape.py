import sys
import getopt
import requests
from bs4 import BeautifulSoup as soup
from markdownify import markdownify
from dateutil.tz import gettz
from datetime import datetime

from utils import blocker


def main(argv):
    year = 2020
    day = 7
    filename = "puzzle.txt"
    try:
        opts, _ = getopt.getopt(argv, "hy:d:o:", ["year=", "day=", "output="])
    except getopt.GetoptError:
        print("scrape.py -y <year> -d <day>")
        sys.exit(2)
    for opt, arg in opts:
        if opt == "-h":
            print("test.py -y <year> -d <day>")
            sys.exit()
        elif opt in ("-y", "--year"):
            year = int(arg)
        elif opt in ("-d", "--day"):
            day = int(arg)
        elif opt in ("-o", "--output"):
            filename = arg
    AOC_TZ = gettz("America/New_York")
    now = datetime.now(tz=AOC_TZ)
    if year == now.year and day > now.day:
        blocker()
    body = requests.get(
        f"https://adventofcode.com/{year}/day/{day}",
        cookies={
            "session": "53616c7465645f5f5d35af4c2ac952f04d3168219c7f761b9975bbeb44b44f665617f224428db73dcbd41ca1707ea0b2"
        },
    )
    html = soup(body.content, "html.parser")
    articles = html.body.main.find_all("article")
    with open(filename, "w+") as f:
        [f.write(markdownify(str(article))) for article in articles]


if __name__ == "__main__":
    main(sys.argv[1:])
