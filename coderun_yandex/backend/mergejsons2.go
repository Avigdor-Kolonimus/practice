package backend

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type JSON struct {
	Offers []Offer `json:"offers"`
}

func (js *JSON) ToJson() string {
	res := strings.Builder{}
	res.WriteString("{\"offers\":[")
	for i, v := range js.Offers {
		res.WriteString(v.ToJson())
		if i != len(js.Offers)-1 {
			res.WriteString(", ")
		}
	}
	res.WriteString("]}")
	return res.String()
}

type Offer struct {
	OfferId   string `json:"offer_id"`
	MarketSku int    `json:"market_sku"`
	Price     int    `json:"price"`
}

func (offer *Offer) ToJson() string {
	return fmt.Sprintf(`{"market_sku":%d,"offer_id":"%s","price":%d}`, offer.MarketSku, offer.OfferId, offer.Price)
}

func validateMergeJSONs2NInput(n int) bool {
	return n >= 1 && n <= 200
}

func validateMergeJSONs2MInput(n int) bool {
	return n >= 1 && n <= 40_000
}

// https://coderun.yandex.ru/selections/backend/problems/merge-jsons-2
// MergeJSONs2 - problem 50
func MergeJSONs2() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	firstLine := mustReadIntArray(reader, 2)
	if !validateMergeJSONs2NInput(firstLine[0]) {
		panic("number N out of range")
	}
	if !validateMergeJSONs2MInput(firstLine[1]) {
		panic("number M out of range")
	}
	n, m := firstLine[0], firstLine[1]

	J := make([]JSON, n)
	for i := range n {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		json.Unmarshal([]byte(line), &J[i])
	}

	res := JSON{Offers: []Offer{}}
	for _, v := range J {
		for _, offer := range v.Offers {
			if len(res.Offers) < m {
				res.Offers = append(res.Offers, offer)
			}
		}
	}

	writer.WriteString(res.ToJson())
	writer.WriteByte('\n')
}
