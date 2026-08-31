import Foundation

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/route-reduction
// RouteReduction - problem 23
func routeReduction() {
    var result = [(direction: String, distance: Int)]()

    while let line = readLine(), !line.isEmpty {
        let parts = line.split(separator: " ")
        let direction = String(parts[0])
        let distance = Int(parts[1])!

        if let last = result.last, last.direction == direction {
            result[result.count - 1].distance += distance
            continue
        }

        if let last = result.last,
           last.direction == oppositeDirection(direction) {

            if last.distance > distance {
                result[result.count - 1].distance -= distance
            } else if last.distance < distance {
                result.removeLast()
                result.append((
                    direction: direction,
                    distance: distance - last.distance
                ))
            } else {
                result.removeLast()
            }

            mergeLastIfNeeded(&result)

            continue
        }

        result.append((direction, distance))
    }

    for item in result {
        print("\(item.direction) \(item.distance)")
    }
}

func oppositeDirection(_ direction: String) -> String {
    switch direction {
    case "LEFT":
        return "RIGHT"
    case "RIGHT":
        return "LEFT"
    case "TOP":
        return "BOTTOM"
    case "BOTTOM":
        return "TOP"
    default:
        return ""
    }
}

func mergeLastIfNeeded(
    _ result: inout [(direction: String, distance: Int)]
) {
    guard result.count >= 2 else {
        return
    }

    let lastIndex = result.count - 1
    let previousIndex = result.count - 2

    if result[previousIndex].direction == result[lastIndex].direction {
        result[previousIndex].distance += result[lastIndex].distance
        result.removeLast()
    }
}

routeReduction()