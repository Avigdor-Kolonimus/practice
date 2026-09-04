import Foundation

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/rle-test
// RleTest - problem 32
func rletest() {
    let s: String = readLine()!

    let result = s.components(separatedBy: .letters).reduce(0) { 
        sum, part in
        sum + (Int(part) ?? 1) } - 1

    print(result)
}

rletest() 