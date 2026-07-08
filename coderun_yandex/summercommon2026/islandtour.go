package summercommon2026

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// API structures

type Tour struct {
	StartIsland int `json:"start_island"`
	MaxStamps   int `json:"max_stamps"`
}

type Route struct {
	From int `json:"from"`
	To   int `json:"to"`
}

type RoutesResponse struct {
	Routes []Route `json:"routes"`
}

type Stamp struct {
	StampID  int   `json:"stamp_id"`
	IslandID int   `json:"island_id"`
	Rarity   int   `json:"rarity"`
	Requires []int `json:"requires"`
}

type StampsResponse struct {
	Stamps []Stamp `json:"stamps"`
}

// HTTP request helpers
func getTour(baseURL, tourID string) (Tour, error) {
	resp, err := http.Get(fmt.Sprintf("%s/tour/%s", baseURL, tourID))
	if err != nil {
		return Tour{}, err
	}
	defer resp.Body.Close()

	var t Tour
	err = json.NewDecoder(resp.Body).Decode(&t)

	return t, err
}

func getRoutes(baseURL, tourID string) ([]Route, error) {
	resp, err := http.Get(fmt.Sprintf("%s/routes?tour_id=%s", baseURL, tourID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r RoutesResponse
	err = json.NewDecoder(resp.Body).Decode(&r)

	return r.Routes, err
}

func getStamps(baseURL, tourID string) ([]Stamp, error) {
	resp, err := http.Get(fmt.Sprintf("%s/stamps?tour_id=%s", baseURL, tourID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var s StampsResponse
	err = json.NewDecoder(resp.Body).Decode(&s)

	return s.Stamps, err
}

// https://coderun.yandex.ru/seasons/2026-summer/tracks/common/problem/island-tour
// IslandTour - problem 8
func IslandTour() {
	in := bufio.NewReader(os.Stdin)
	baseURL, _ := in.ReadString('\n')
	tourID, _ := in.ReadString('\n')
	baseURL = strings.TrimSpace(baseURL)
	tourID = strings.TrimSpace(tourID)

	if baseURL == "" || tourID == "" {
		fmt.Println(0)

		return
	}

	// 1. Fetch data from the API
	tour, err := getTour(baseURL, tourID)
	if err != nil {
		fmt.Println(0)

		return
	}

	routes, err := getRoutes(baseURL, tourID)
	if err != nil {
		fmt.Println(0)

		return
	}

	stamps, err := getStamps(baseURL, tourID)
	if err != nil {
		fmt.Println(0)

		return
	}

	// 2. Find reachable islands using BFS
	adjIslands := make(map[int][]int)
	for _, r := range routes {
		adjIslands[r.From] = append(adjIslands[r.From], r.To)
	}

	reachableIslands := make(map[int]bool)
	queue := []int{tour.StartIsland}
	reachableIslands[tour.StartIsland] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, next := range adjIslands[curr] {
			if !reachableIslands[next] {
				reachableIslands[next] = true
				queue = append(queue, next)
			}
		}
	}

	// 3. Filter stamps by island reachability
	stampMap := make(map[int]Stamp)
	for _, s := range stamps {
		stampMap[s.StampID] = s
	}

	// Verify if the stamp and all its prerequisites are reachable
	validStamps := make(map[int]bool)
	for _, s := range stamps {
		curr := s
		possible := true
		for {
			if !reachableIslands[curr.IslandID] {
				possible = false
				break
			}
			if len(curr.Requires) == 0 {
				break
			}
			parentID := curr.Requires[0]
			var ok bool
			curr, ok = stampMap[parentID]
			if !ok {
				possible = false
				break
			}
		}

		if possible {
			validStamps[s.StampID] = true
		}
	}

	// 4. Build the dependency tree (with a dummy root ID = 0)
	// The dummy root has rarity = 0 and is always "selected" (hence capacity + 1)
	type Node struct {
		id       int
		rarity   int
		children []int
	}

	nodes := make(map[int]*Node)
	nodes[0] = &Node{id: 0, rarity: 0}

	for _, s := range stamps {
		if !validStamps[s.StampID] {
			continue
		}
		nodes[s.StampID] = &Node{id: s.StampID, rarity: s.Rarity}
	}

	for _, s := range stamps {
		if !validStamps[s.StampID] {
			continue
		}

		if len(s.Requires) == 0 {
			nodes[0].children = append(nodes[0].children, s.StampID)
		} else {
			parentID := s.Requires[0]
			if validStamps[parentID] {
				nodes[parentID].children = append(nodes[parentID].children, s.StampID)
			}
		}
	}

	k := tour.MaxStamps
	dp := make(map[int][]int)
	choices := make(map[int][][]int)

	var dfs func(int)
	dfs = func(u int) {
		node := nodes[u]
		res := make([]int, k+2)
		for i := range res {
			res[i] = -1 // -1 signifies an unreachable state
		}
		res[1] = node.rarity // The node itself must be taken (weight 1)

		childChoices := make([][]int, len(node.children))

		for childIdx, child := range node.children {
			dfs(child)
			childDp := dp[child]

			nextRes := make([]int, k+2)
			for i := range nextRes {
				nextRes[i] = -1
			}

			childChoices[childIdx] = make([]int, k+2)

			for i := 1; i <= k+1; i++ {
				if res[i] == -1 {
					continue
				}

				// Option 1: Take nothing from this child
				if res[i] > nextRes[i] {
					nextRes[i] = res[i]
					childChoices[childIdx][i] = 0
				}

				// Option 2: Take j items from the child's subtree (j >= 1, since the child must be picked)
				for j := 1; i+j <= k+1; j++ {
					if childDp[j] == -1 {
						continue
					}
					val := res[i] + childDp[j]
					if val > nextRes[i+j] {
						nextRes[i+j] = val
						childChoices[childIdx][i+j] = j
					}
				}
			}
			res = nextRes
		}
		dp[u] = res
		choices[u] = childChoices
	}

	dfs(0)

	// Find the best capacity (the dummy root weighs 1, so we scan from 1 to k+1)
	bestW := 1
	maxScore := 0
	for w := 1; w <= k+1; w++ {
		if dp[0][w] > maxScore {
			maxScore = dp[0][w]
			bestW = w
		}
	}

	if maxScore == 0 || bestW <= 1 {
		fmt.Println(0)

		return
	}

	// 5. Reconstruct the selected stamps
	var picked []int
	var recoverAns func(int, int)
	recoverAns = func(u int, w int) {
		if u != 0 {
			picked = append(picked, u)
		}
		node := nodes[u]
		currW := w

		// Iterate backwards through children since the DP was updated sequentially
		for i := len(node.children) - 1; i >= 0; i-- {
			child := node.children[i]
			taken := choices[u][i][currW]
			if taken > 0 {
				recoverAns(child, taken)
				currW -= taken
			}
		}
	}

	recoverAns(0, bestW)

	// Topological sort for the output (ensuring prerequisites always appear first)
	// Build a dependency graph containing only the selected stamps
	pickedMap := make(map[int]bool)
	for _, id := range picked {
		pickedMap[id] = true
	}

	inDegree := make(map[int]int)
	adjPicked := make(map[int][]int)
	for _, id := range picked {
		s := stampMap[id]
		if len(s.Requires) > 0 && pickedMap[s.Requires[0]] {
			parent := s.Requires[0]
			adjPicked[parent] = append(adjPicked[parent], id)
			inDegree[id]++
		}
	}

	var order []int
	var q []int
	for _, id := range picked {
		if inDegree[id] == 0 {
			q = append(q, id)
		}
	}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		order = append(order, curr)
		for _, next := range adjPicked[curr] {
			inDegree[next]--
			if inDegree[next] == 0 {
				q = append(q, next)
			}
		}
	}

	// Print results
	fmt.Println(len(order))
	strIDs := make([]string, len(order))
	for i, id := range order {
		strIDs[i] = strconv.Itoa(id)
	}
	fmt.Println(strings.Join(strIDs, " "))
}
