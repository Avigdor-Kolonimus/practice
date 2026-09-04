import Foundation

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/mark
// Mark - problem 29
func mark() {
    let numbers = readLine()!.split(separator: " ").map { Int64($0)! }

    let sorted = numbers.sorted()

    print(sorted[1])
}

mark()