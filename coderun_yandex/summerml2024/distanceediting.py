import sys

# https://coderun.yandex.ru/selections/2024-summer-ml/problems/distance-editing
# DistanceEditing - problem 22
def main():
    n, m = map(int, input().split())    
    s, t = input(), input()
    I, D, S = map(int, input().split())

    matrix = [[0] * (m + 1) for _ in range(n + 1)]

    for i in range(1, n + 1):  matrix[i][0] = i * D
    for j in range(1, m + 1):  matrix[0][j] = j * I

    for i in range(1, n + 1):
        for j in range(1, m + 1):
            if s[i - 1] == t[j -1]: matrix[i][j] = matrix[i - 1][j -1 ]
            else: matrix[i][j] = min(matrix[i - 1][j] + D, matrix[i - 1][j - 1] + S, matrix[i][j - 1] + I)
    
    print(matrix[-1][-1])

if __name__ == '__main__':
    main()