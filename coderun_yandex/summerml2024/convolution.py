import sys

# https://coderun.yandex.ru/selections/2024-summer-ml/problems/convolution
# Convolution - problem 24
def main():
    n, m = map(int, input().split())

    A = []
    for _ in range(n):
        row = list(map(int, input().split()))
        A.append(row)

    k = int(input())

    B = []
    for _ in range(k):
        row = list(map(int, input().split()))
        B.append(row)

    res_n = n - k + 1
    res_m = m - k + 1

    C = [[0 for _ in range(res_m)] for _ in range(res_n)]
    for i in range(res_n):
        for j in range(res_m):
            total = 0
            for t in range(k):
                for l in range(k):
                    total += A[i + t][j + l] * B[t][l]
            C[i][j] = total

    for row in C:
        print(' '.join(map(str, row)))


if __name__ == '__main__':
    main()
