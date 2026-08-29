import Foundation

func convertToSeconds(_ time: String) -> Int {
    let h = Int(time.prefix(2)) ?? 0
    let m = Int(time.dropFirst(3).prefix(2)) ?? 0
    let s = Int(time.suffix(2)) ?? 0
    return h * 3600 + m * 60 + s
}

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/log-without-dates
// LogWithoutDates - problem 6
func main(){
    let n = Int(readLine() ?? "0") ?? 0
    var days = 1
    var previousSeconds = -1

    for _ in 0..<n {
        let time = readLine() ?? "00:00:00"
        let currentSeconds = convertToSeconds(time)
    
        if currentSeconds <= previousSeconds {
            days += 1
        }
        previousSeconds = currentSeconds
    }

    print(days)
}

main()