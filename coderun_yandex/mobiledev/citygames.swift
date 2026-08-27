import Foundation

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/city-games
// CityGames - problem 5
func parse(_ input: String) -> Set<String> {
    var result = Set<String>()

    let trimmed = input.trimmingCharacters(in: .whitespacesAndNewlines)
    let text = String(trimmed.dropLast())

    for pair in text.split(separator: ",") {
        let names = pair.split(separator: " ").map(String.init)

        if names.count == 2 {
            let sortedNames = names.sorted()
            let key = sortedNames[0] + "|" + sortedNames[1]
            result.insert(key)
        }
    }

    return result
}

let line1 = readLine()!
let line2 = readLine()!

let set1 = parse(line1)
let set2 = parse(line2)

print(set1 == set2 ? "YES" : "NO")