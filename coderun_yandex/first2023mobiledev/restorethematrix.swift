import Foundation

func readIntegers() -> [Int] {
    let line = readLine()!.split(separator: " ")
    return line.map { Int($0)! }
}

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/restore-the-matrix
// RestoreTheMatrix - problem 21
func restoreTheMatrix() {   
    let n = readIntegers()[0]
    var a = Array(repeating: Array<Int>(), count: n)
    
    var unused = Set(1 ... n * n)
    for i in 0 ..< n {
        a[i] = readIntegers()
        a[i].forEach { unused.remove($0) }
    }

    for i in 0 ..< n {
        for j in 0 ..< n {
            if a[i][j] == 0, let currentValue = unused.first {
                print(currentValue, terminator: " ")
                unused.remove(currentValue)
            } else {
                print(a[i][j], terminator: " ")
            }
        }
        print("\n", terminator: "")
    }
}

restoreTheMatrix()