import numpy as np

# https://coderun.yandex.ru/selections/first-2023-ml/problems/warm-up
# WarmUp - problem 17
def main():
    train = np.loadtxt("train.tsv")
    test = np.loadtxt("test.tsv")

    # first 100 column
    X = train[:, :100]

    # 101 column — target
    y = train[:, 100]

    X = np.column_stack((X, np.ones(len(X))))

    # X * w = y
    w, _, _, _ = np.linalg.lstsq(X, y, rcond=None)

    test_X = np.column_stack((test, np.ones(len(test))))
    predictions = test_X @ w

    # answer.tsv
    with open("answer.tsv", "w") as out:
        for value in predictions:
            out.write(f"{value:.8f}\n")


if __name__ == "__main__":
    main()