import Foundation

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/lzw-mob
// LzwMob - problem 16
class TrieNode {
    var zero: TrieNode? = nil  
    var one:  TrieNode? = nil  
    var isEndOfBlock: Bool = false
}

func insert(_ root: TrieNode, _ chars: [Character], _ start: Int, _ length: Int) {
    var current = root
    for i in start..<(start + length) {
        let char = chars[i]
        if char == "1" {
            let next = current.one ?? TrieNode()
            current.one = next
            current = next
        } else {
            let next = current.zero ?? TrieNode()
            current.zero = next
            current = next
        }
    }
    current.isEndOfBlock = true
}

func findLongestMatch(_ chars: [Character], _ startPos: Int, _ root: TrieNode) -> Int {
    var currentNode = root
    var candidateLen = 0
    var longestMatchLen = 0
    
    var idx = startPos
    while idx < chars.count {
        let char = chars[idx]
        switch char {
        case "0":
            guard let next = currentNode.zero else { return longestMatchLen }
            currentNode = next
        case "1":
            guard let next = currentNode.one else { return longestMatchLen }
            currentNode = next
        default:
            print("Error: error character")
            exit(0)
        }
        candidateLen += 1
        idx += 1
        
        if currentNode.isEndOfBlock {
            longestMatchLen = candidateLen
        }
    }
    
    return longestMatchLen
}

guard let line = readLine() else {
    print("Error: zero input")
    exit(0)
}

let chars = Array(line.trimmingCharacters(in: .whitespacesAndNewlines))
guard !chars.isEmpty else {
    print("Error: empty string")
    exit(0)
}

let root = TrieNode()
var blockRanges: [(start: Int, length: Int)] = []
var position = 0


insert(root, chars, 0, 1)
blockRanges.append((start: 0, length: 1))
position = 1

while position < chars.count {
    let matchLen = findLongestMatch(chars, position, root)
    let remainingLen = chars.count - position
    
    let blockLength: Int
    if matchLen == remainingLen {
        blockLength = matchLen
    } else {
        blockLength = matchLen + 1 
    }
    
    insert(root, chars, position, blockLength)
    blockRanges.append((start: position, length: blockLength))
    position += blockLength
}

let result = blockRanges.map { range in
    String(chars[range.start..<(range.start + range.length)])
}.joined(separator: " ")

print(result)