import Foundation

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/palindroming-check
// PalindromingCheck - problem 30
func palindromingCheck() {
  let s = Array(readLine()!.lowercased())

  var left = 0
  var right = s.count - 1

  var isPalindrome = true
  while left < right {
    while left < right && s[left] == " " {
        left += 1
    }

    while left < right && s[right] == " " {
        right -= 1
    }

    if s[left] != s[right] {
        isPalindrome = false
        break
    }

    left += 1
    right -= 1
  }

  print(isPalindrome ? "It is a palindrome" : "It is not a palindrome")
}

palindromingCheck() 