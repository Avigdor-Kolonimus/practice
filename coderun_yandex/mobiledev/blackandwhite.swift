import Foundation

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/black-and-white
// BlackAndWhite - problem 21
func blackAndWhite() {
    let input1 = readLine()!.split(separator: " ").map { Int(String($0))! }
    let n = input1.first!
    let m = input1.last!
    
    var rows = [String]()
    for _ in 0..<n {
        let row = readLine()!
        rows.append(row)
    }

    var result = 0
    for row in rows {
        var isInside = false
        
        for char in row {
            switch char {
            case "/", "\\":
                if isInside {
                    result += 1
                }
                isInside.toggle()
            case ".":
                if isInside {
                    result += 1
                }
            default:
                continue
            }
        }
    }
    
    print(result)
}

blackAndWhite()