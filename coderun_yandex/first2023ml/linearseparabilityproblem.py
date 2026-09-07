import sys

# https://coderun.yandex.ru/selections/first-2023-ml/problems/linear-separability-problem
# LinearSeparabilityProblem - problem 9
def main():
    input = sys.stdin.buffer.readline

    n, m = map(int, input().split())

    x = []
    y = []

    for _ in range(n):
        data = list(map(float, input().split()))
        x.append(data[:m])
        y.append(data[m])

    a = [0.0] * m

    while True:
        changed = False

        for i in range(n):
            dot = 0.0

            for j in range(m):
                dot += a[j] * x[i][j]

            if y[i] * dot <= 0:
                changed = True

                for j in range(m):
                    a[j] += y[i] * x[i][j]

        if not changed:
            break

    print(*a)


if __name__ == '__main__':
    main()