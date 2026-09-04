import Foundation

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/bug-in-library
// BugInLibrary - problem 14
func bugInLibrary() {
    var strings = [String]()
    while let string = readLine() {
        if string == "" {
            break
        }
        strings.append(string)
    }
    var result = Array(repeating: "", count: strings.count)
    
    for str in strings {
        let arr = Array(str)
        var num = ""
        var s = ""
        for ch in arr {
            if ch.isNumber {
                num += String(ch)
            } else {
                s += String(ch)
            }
        }
        result[Int(num)! - 1] = s
    }
    result.forEach { print($0) }
}

bugInLibrary()