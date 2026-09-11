import Foundation

func binSearch(_ arr: [Int], target: Int) -> Int {
  var left = 0, right = arr.count - 1
  while right - left > 1 {
    let mid = (right + left) / 2
    if arr[mid] == target {
      return target
    } else if arr[mid] < target {
      left = mid
    } else {
      right = mid
    }
  }

  if abs(arr[left] - target) <= abs(arr[right] - target) {
    return arr[left]
  } else {
    return arr[right]
  }
}

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/normalization-of-indicators
// NormalizationOfIndicators - problem 21
func normalizationofindicators() {
    let n = Int(readLine()!)!
    let newAns = readLine()!.split(separator: " ")
    .map { Int($0)! }
    .sorted()
    let m = Int(readLine()!)!
    var oldAns = [Int]()
    for _ in 0..<m { oldAns.append(Int(readLine()!)!) }
    for ans in oldAns {
        let result = binSearch(newAns, target: ans)
        print(result)
    }
}

normalizationofindicators()
