import sys
from itertools import permutations

def minimal_deg(seq):
    arr = seq[:]
    for k in range(len(seq)):
        if all(x == arr[0] for x in arr):
            return k
        arr = [arr[i+1] - arr[i] for i in range(len(arr)-1)]
    return len(seq)-1

# https://coderun.yandex.ru/selections/first-2023-data-analytics/problems/restore-polynomial
# RestorePolynomial - problem 27
def main():
    data = sys.stdin.read().split()
    t = int(data[0])
    ptr = 1
    out = []

    for _ in range(t):
        n = int(data[ptr]); ptr += 1
        a = list(map(int, data[ptr:ptr+n]))
        ptr += n

        best = n-1
        for perm in permutations(a):
            d = minimal_deg(perm)
            if d < best:
                best = d
                if best == 0:
                    break 
        out.append(str(best))

    sys.stdout.write("\n".join(out))

if __name__ == "__main__":
    main()