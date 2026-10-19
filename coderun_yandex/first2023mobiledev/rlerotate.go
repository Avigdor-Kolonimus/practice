package first2023mobiledev

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Segment struct {
	length int64
	ch     byte
}

func addSegment(segments *[]Segment, length int64, ch byte) {
	if length == 0 {
		return
	}

	arr := *segments
	if len(arr) > 0 && arr[len(arr)-1].ch == ch {
		arr[len(arr)-1].length += length
	} else {
		*segments = append(*segments, Segment{length, ch})
	}
}

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/rle-rotate
// RleRotate - problem 23
func RleRotate() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	n, _ := strconv.Atoi(parts[0])
	k, _ := strconv.ParseInt(parts[1], 10, 64)

	line, _ = reader.ReadString('\n')
	s := strings.TrimSpace(line)

	segments := make([]Segment, 0, n)
	var length int64

	for i := 0; i < len(s); {
		var count int64
		hasCount := false

		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			hasCount = true
			count = count*10 + int64(s[i]-'0')
			i++
		}

		if i >= len(s) {
			break
		}

		ch := s[i]
		i++

		if !hasCount {
			count = 1
		}

		segments = append(segments, Segment{count, ch})

		if length <= k {
			if count > k-length {
				length = k + 1
			} else {
				length += count
			}
		}
	}

	shift := k
	if length <= k {
		shift = k % length
	}

	prefix := make([]Segment, 0, len(segments)+1)
	suffix := make([]Segment, 0, len(segments)+1)

	remaining := shift
	cut := false

	for _, seg := range segments {
		if cut {
			addSegment(&suffix, seg.length, seg.ch)
			continue
		}

		if remaining >= seg.length {
			addSegment(&prefix, seg.length, seg.ch)
			remaining -= seg.length
			continue
		}

		if remaining > 0 {
			addSegment(&prefix, remaining, seg.ch)
			addSegment(&suffix, seg.length-remaining, seg.ch)
		} else {
			addSegment(&suffix, seg.length, seg.ch)
		}

		cut = true
	}

	result := make([]Segment, 0, len(segments)+2)

	for _, seg := range suffix {
		addSegment(&result, seg.length, seg.ch)
	}
	for _, seg := range prefix {
		addSegment(&result, seg.length, seg.ch)
	}

	for _, seg := range result {
		if seg.length > 1 {
			writer.WriteString(strconv.FormatInt(seg.length, 10))
		}
		writer.WriteByte(seg.ch)
	}
	writer.WriteByte('\n')
}
