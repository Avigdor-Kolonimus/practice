import Foundation

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/report-restoration
// ReportRestoration - problem 9
func reportRestoration() {
    let input1 = readLine()!.split(separator: " ").map { Int(String($0))! }
    let pagesCount = input1[0]
    let reportCount = input1[1]
    var pages: Set<String> = Set(readLine()!.split(separator: " ").map { String($0) })
    var dict = [String: [String]]()
    let main = readLine()!
   
    for _ in 1...reportCount {
        let line = readLine()!.split(separator: " ").map { String($0) }
        dict[line[0], default: []].append(line[1])
    }

    var visited = Set<String>()
    var stack = [String]()
    stack.append(main)
    while !stack.isEmpty {
        let currentPage = stack.removeLast()
        if visited.contains(currentPage) {
            continue
        }
        visited.insert(currentPage)
        if pages.contains(currentPage) {
            pages.remove(currentPage)
        }
        if let nextPages = dict[currentPage] {
            stack.append(contentsOf: nextPages)
        }
        
    }
    
    print(pages.count)
}

reportRestoration()