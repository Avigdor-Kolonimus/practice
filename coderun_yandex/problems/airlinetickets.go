package problems

// https://coderun.yandex.ru/problem/airline-tickets
// AirlineTickets - problem 754
func AirlineTickets() []string {
	flights := map[string][]string{
		"A": {"B", "D"},
		"B": {"C", "N", "Z"},
		"D": {"E", "F"},
		"F": {"S"},
	}

	fetchFlights := func(from string) []string {
		return flights[from]
	}

	return FindPath("A", "N", fetchFlights)
}
func FindPath(from string, to string, fetchFlights func(string) []string) []string {
	path := []string{from}
	stack := [][]string{fetchFlights(from)}

	for len(stack) > 0 {
		if path[len(path)-1] == to {
			return path
		}

		nodes := stack[len(stack)-1]

		if len(nodes) == 0 {
			path = path[:len(path)-1]
			stack = stack[:len(stack)-1]
			continue
		}

		node := nodes[len(nodes)-1]
		stack[len(stack)-1] = nodes[:len(nodes)-1]

		path = append(path, node)
		stack = append(stack, fetchFlights(node))
	}

	return []string{}
}
