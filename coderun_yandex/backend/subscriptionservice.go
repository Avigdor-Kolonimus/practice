package backend

import (
	"bufio"
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

type Subscription struct {
	triggers  [][]string
	shipments [][]string
}

type Message struct {
	TraceID string                 `json:"trace_id"`
	Offer   map[string]interface{} `json:"offer"`
}

func splitPath(path string) []string {
	return strings.Split(path, ".")
}

// merge combines src with dst and returns the changed leaf-paths.
func merge(dst, src map[string]interface{}, prefix []string) []string {
	var changed []string

	for key, newValue := range src {
		path := append(prefix, key)
		oldValue, exists := dst[key]

		if !exists {
			dst[key] = newValue
			changed = append(changed, collectLeafPaths(newValue, path)...)

			continue
		}

		oldObj, oldIsObj := oldValue.(map[string]interface{})
		newObj, newIsObj := newValue.(map[string]interface{})

		if oldIsObj && newIsObj {
			changed = append(
				changed,
				merge(oldObj, newObj, path)...,
			)
			continue
		}

		if !jsonEqual(oldValue, newValue) {
			dst[key] = newValue
			changed = append(
				changed,
				strings.Join(path, "."),
			)
		}
	}

	return changed
}

func collectLeafPaths(value interface{}, prefix []string) []string {
	obj, ok := value.(map[string]interface{})
	if !ok {
		return []string{strings.Join(prefix, ".")}
	}

	var result []string
	for key, child := range obj {
		result = append(
			result,
			collectLeafPaths(
				child,
				append(prefix, key),
			)...,
		)
	}

	return result
}

func jsonEqual(a, b interface{}) bool {
	ba, _ := json.Marshal(a)
	bb, _ := json.Marshal(b)

	return string(ba) == string(bb)
}

// We check whether the changedPath is nested within thetriggerPath.
//
// trigger:  partner_content
// changed:  partner_content.title
//
// trigger:  partner_content.title
// changed:  partner_content.title
func isInside(
	changedPath []string,
	triggerPath []string,
) bool {
	if len(changedPath) < len(triggerPath) {
		return false
	}

	for i := range triggerPath {
		if changedPath[i] != triggerPath[i] {
			return false
		}
	}

	return true
}

func subscriptionTriggered(sub Subscription, changedPaths []string) bool {
	for _, changed := range changedPaths {
		changedParts := splitPath(changed)

		for _, trigger := range sub.triggers {
			if isInside(changedParts, trigger) {
				return true
			}
		}
	}

	return false
}

// Retrieves a value via a JSON path
func getPath(obj map[string]interface{}, path []string) (interface{}, bool) {
	var current interface{} = obj
	for _, key := range path {
		m, ok := current.(map[string]interface{})
		if !ok {
			return nil, false
		}

		current, ok = m[key]
		if !ok {
			return nil, false
		}
	}

	return current, true
}

// Copies a field from src to dst using a JSON path
func copyPath(src, dst map[string]interface{}, path []string) {
	value, ok := getPath(src, path)
	if !ok {
		return
	}

	current := dst
	for i := 0; i < len(path)-1; i++ {
		key := path[i]

		next, ok := current[key].(map[string]interface{})
		if !ok {
			next = make(map[string]interface{})
			current[key] = next
		}

		current = next
	}

	current[path[len(path)-1]] = value
}

// Создаём offer для конкретной подписки.
// В него попадают только trigger + shipment поля.
func buildOffer(offer map[string]interface{}, sub Subscription) map[string]interface{} {
	result := make(map[string]interface{})

	// offer.id mandatory
	if id, ok := offer["id"]; ok {
		result["id"] = id
	}

	for _, path := range sub.triggers {
		copyPath(offer, result, path)
	}

	for _, path := range sub.shipments {
		copyPath(offer, result, path)
	}

	return result
}

// https://coderun.yandex.ru/selections/backend/problems/subscription-service
// SubscriptionService - problem 25
func SubscriptionService() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	line, _ := reader.ReadString('\n')
	fields := strings.Fields(line)

	n, _ := strconv.Atoi(fields[0])
	m, _ := strconv.Atoi(fields[1])

	// Subscription input
	subscriptions := make([]Subscription, n)
	for i := range n {
		line, _ = reader.ReadString('\n')
		fields = strings.Fields(line)

		pos := 0

		a, _ := strconv.Atoi(fields[pos])
		pos++

		b, _ := strconv.Atoi(fields[pos])
		pos++

		sub := Subscription{
			triggers:  make([][]string, 0, a),
			shipments: make([][]string, 0, b),
		}

		for j := 0; j < a; j++ {
			sub.triggers = append(
				sub.triggers,
				splitPath(fields[pos]),
			)
			pos++
		}

		for j := 0; j < b; j++ {
			sub.shipments = append(
				sub.shipments,
				splitPath(fields[pos]),
			)
			pos++
		}

		subscriptions[i] = sub
	}

	// Current state of all offers.
	offers := make(map[string]map[string]interface{})

	// Update input
	for range m {
		line, _ = reader.ReadString('\n')

		var message Message

		json.Unmarshal(
			[]byte(strings.TrimSpace(line)),
			&message,
		)

		offerID := message.Offer["id"].(string)

		offer, exists := offers[offerID]
		if !exists {
			offer = make(map[string]interface{})
			offers[offerID] = offer
		}

		// Apply partial update.
		changedPaths := merge(
			offer,
			message.Offer,
			nil,
		)

		// Subscribers are processed in their original order.
		for _, sub := range subscriptions {
			if !subscriptionTriggered(
				sub,
				changedPaths,
			) {
				continue
			}

			result := Message{
				TraceID: message.TraceID,
				Offer: buildOffer(
					offer,
					sub,
				),
			}

			data, _ := json.Marshal(result)

			writer.WriteString(string(data))
			writer.WriteByte('\n')
		}
	}
}
