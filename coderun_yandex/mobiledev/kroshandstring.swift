import Foundation

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/krosh-and-string
// KroshAndString - problem 17
func main(){
    let _ = readLine()
    var input = readLine()!
    var array = [String]()

    input.forEach { character in
        if !array.isEmpty && String(character) == array.last!  {
            array.removeLast()
        } else {
            array.append(String(character))
        }
    }

    print(array.isEmpty ? "1" : "0")
}

main()
