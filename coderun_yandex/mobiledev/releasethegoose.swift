import Foundation

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/release-the-goose
// ReleaseTheGoose - problem 2
func main() {
    let (n, k) = {
        let components = readLine()!.components(separatedBy: " ").compactMap(Int.init)
        return (components[0], components[1])
    }()
   
    var data = [String: Int]()
    var site: String, description: String
    var result = Set<String>()
    for _ in 0..<n {
        site = readLine()!
        description = readLine()!
        if data[site, default: 0] >= k {
            continue
        }
        
        for word in description.components(separatedBy: " ") where word == "goose" && data[site, default: 0] < k {
            data[site, default: 0] += 1
        }
        
        if data[site, default: 0] >= k {
            result.insert(site)
        }
    }

    print(result.count)
    
    for site in result.sorted() {
        print(site)
    }
}

main()