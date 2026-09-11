import sys

def get_mask(x: int) -> int:
    mask = 0

    if x == 0:
        return 1

    while x > 0:
        digit = x % 10
        mask |= (1 << digit)
        x //= 10

    return mask

# https://coderun.yandex.ru/selections/first-2023-data-analytics/problems/three-numbers
# ThreeNumbers - problem 32
def main() -> None:
    input = sys.stdin.readline

    t = int(input())
    ALL = (1 << 10) - 1

    for _ in range(t):
        n = int(input())
        a = list(map(int, input().split()))

        masks = [get_mask(x) for x in a]

        found = False

        for i in range(n):
            if found:
                break

            for j in range(i + 1, n):
                if found:
                    break

                for k in range(j + 1, n):
                    if (masks[i] | masks[j] | masks[k]) == ALL:
                        print(a[i], a[j], a[k])
                        found = True
                        break

if __name__ == '__main__':
    main()
