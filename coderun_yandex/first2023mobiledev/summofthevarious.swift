import Foundation

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/summ-of-the-various
// SummOfTheVarious - problem 35
func SummOfTheVarious() {
    _ = Int(readLine()!)!
    let numberArray = readLine()!.split(separator: " ").map { Int($0) ?? 0 }

    print(Set(numberArray).reduce(0) { $0 + $1 })
}

SummOfTheVarious()