"""
This example module calculates the number pi using the Leibniz formula
"""

import sys


def leibniz_sum(start: int, end: int) -> float:
    """ ref: https://en.wikipedia.org/wiki/Leibniz_formula_for_pi """
    result = 0.0
    for i in range(start, end):
        result += (4 * (-1) ** i) / (2 * i + 1)

    return result


def main():
    start = int(sys.argv[1])
    end = int(sys.argv[2])
    print(leibniz_sum(start, end))


if __name__ == '__main__':
    main()
