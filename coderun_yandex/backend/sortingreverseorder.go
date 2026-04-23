package backend

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/backend/problems/rocks-and-jewels
// SortingReverseOrder - problem 53
func SortingReverseOrder() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// host input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	host := strings.TrimRight(line, "\r\n")

	// port input
	line, err = reader.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	port, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// a input
	line, err = reader.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	a, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// b input
	line, err = reader.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	b, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%s:%d?a=%d&b=%d", host, port, a, b)

	res, _ := http.Get(url)
	data, _ := ioutil.ReadAll(res.Body)

	var values []int
	json.Unmarshal(data, &values)

	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	for _, n := range values {
		if n > 0 {
			writer.WriteString(strconv.Itoa(n))
			writer.WriteByte('\n')
		}
	}
}
