import sys
from bisect import bisect_right
from itertools import accumulate

# https://coderun.yandex.ru/selections/2024-summer-ml/problems/fuzzies
# Fuzzies - problem 27
def main():
    tokens = sys.stdin.buffer.read().split()
    count = int(tokens[0])
    budget = int(tokens[1])
    times = sorted(map(int, tokens[2:2 + count]))

    answer = bisect_right(list(accumulate(times)), budget)
    sys.stdout.write(str(answer) + "\n")


if __name__ == '__main__':
    main()