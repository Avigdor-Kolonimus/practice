import sys
import math

# https://coderun.yandex.ru/selections/first-2023-ml/problems/happiness
# Happiness - problem 28
def main():
    N, M = map(int, input().split())

    A = []
    B = []
    for _ in range(0, N):
        s = input().split()
        A.append(float(s[-2]))
        B.append(float(s[-1]))
    
    sum_A = 0
    for a in A:
        sum_A += a
    
    sum_B = 0
    for b in B:
        sum_B += b

    sum_AB = 0
    for a, b in zip(A, B):
        sum_AB += b*math.log(1+a/b)
    
    p = [sum_A/sum_B, math.exp(sum_AB/sum_B) - 1, sum_A/sum_B]

    print(' '.join([str(e) for e in p]))

if __name__ == '__main__':
    main()
