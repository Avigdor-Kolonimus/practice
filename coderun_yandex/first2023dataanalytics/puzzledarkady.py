import sys

# https://coderun.yandex.ru/selections/first-2023-data-analytics/problems/puzzled-arkady
# PuzzledArkady - problem 48
def main():
    p = float(input())
    print(f'{1.0 / (p ** 2 * (1 - p)):.4f}')


if __name__ == '__main__':
    main()