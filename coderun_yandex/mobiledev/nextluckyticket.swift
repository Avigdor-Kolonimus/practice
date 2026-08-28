import Foundation

func digitSum(_ number: Int) -> Int {
    number / 100 +
    number / 10 % 10 +
    number % 10
}

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/next-lucky-ticket
// NextLuckyTicket - problem 3
func main() {
    var number = Int(readLine()!)!
    repeat {
        number += 1
    } while digitSum(number / 1000) != digitSum(number % 1000)

    print(number)
}

main()