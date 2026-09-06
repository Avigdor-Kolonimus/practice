import sys

# https://coderun.yandex.ru/selections/2024-summer-ml/problems/1984
# 1984 - problem 15
def main():
    n, m = list(map(int, input().split()))
    stopwords = [input() for i in range(n)]
    messages = [input() for i in range(m)]

    for message in messages:
        flag = False
        for stopword in stopwords:
            if stopword in message:
                flag = True
                break
        
        if flag:
            print('DELETE')
        else:
            print('KEEP')

if __name__ == '__main__':
    main()