import Foundation

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/simple-suggest
// SimpleSuggest - problem 25
func simplesuggest() {
    let n = Int(readLine()!)!
    var wc = [String: Int]() 

    for _ in 0..<n {
        let word = Array(readLine()!)
        let letter = String(word.first!)

        wc[letter] = (wc[letter] ?? 0) + 1
    }

    print(wc.max { $0.value < $1.value }!.key)
}

simplesuggest() 