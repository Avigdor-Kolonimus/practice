import Foundation

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/gradient
// Gradient - problem 22
func gradient() {
    let input = readLine()!.split(separator: " ").map { Int(String($0))! }
    let w = input[0]
    let h = input[1]

    var rowIntensity = Array(repeating: 0, count: h)
    var columnIntensity = Array(repeating: 0, count: w)
    for row in 0..<h {
        for column in 0..<w {
            let top = (row + 1) * (column + 1)
            let bottom = (h - row) * (w - column)
            
            let intensity = top * bottom
            rowIntensity[row] += intensity
            columnIntensity[column] += intensity
        }
    }
    
    print(rowIntensity.map { String($0)}.joined(separator: " "))
    print(columnIntensity.map { String($0)}.joined(separator: " "))
}


gradient()