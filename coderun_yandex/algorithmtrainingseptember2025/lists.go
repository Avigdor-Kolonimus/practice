package algorithmtrainingseptember2025

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
)

type List struct {
	Digits []int
	IsRoot bool
	Root   *List
	Offset int
}

type Figon struct {
	Lists map[[10]byte]*List
	Out   io.Writer
}

func (f *Figon) addNewSubList(command CommandNewSubList) {
	rootList := f.Lists[command.RootListName]

	f.Lists[command.ListName] = &List{
		Digits: []int(nil),
		IsRoot: true,
		Root:   rootList.Root,
		Offset: rootList.Offset + command.From,
	}
}

func (f *Figon) addToList(command CommandAdd) {
	list := f.Lists[command.ListName]
	if !list.IsRoot {
		return
	}

	list.Digits = append(list.Digits, command.Value)
}

func (f *Figon) setInList(command CommandSet) {
	list := f.Lists[command.ListName]

	list.Root.Digits[list.Offset:][command.Index] = command.Value
}

func (f *Figon) writeListIndex(command CommandGet) {
	list := f.Lists[command.ListName]

	digit := list.Root.Digits[list.Offset:][command.Index]

	var buf []byte

	buf = strconv.AppendInt(buf, int64(digit), 10)
	buf = append(buf, '\n')

	_, err := f.Out.Write(buf)
	if err != nil {
		panic(err)
	}
}

type CommandNewList struct {
	ListName [10]byte
	Digits   []int
}

type CommandNewSubList struct {
	ListName     [10]byte
	RootListName [10]byte
	From, To     int
}

type CommandSet struct {
	ListName [10]byte
	Index    int
	Value    int
}

type CommandGet struct {
	ListName [10]byte
	Index    int
}

type CommandAdd struct {
	ListName [10]byte
	Value    int
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/lists
// ListView - problem 10
func ListView() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	data, err := reader.ReadSlice('\n')
	if err != nil {
		panic(err)
	}

	commandsCount := fastAtoi(data[:len(data)-1])

	figon := new(Figon)
	figon.Out = writer
	figon.Lists = make(map[[10]byte]*List, commandsCount)

	for range commandsCount {
		executeCommand(reader, figon)

		_, err := reader.ReadByte()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			panic(err)
		}
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}

func (f *Figon) addNewList(command CommandNewList) {
	list := &List{
		Digits: command.Digits,
		IsRoot: true,
		Offset: 0,
	}

	list.Root = list

	f.Lists[command.ListName] = list
}

func executeCommand(in *bufio.Reader, figon *Figon) {
	firstBytes, err := in.Peek(1)
	if err != nil {
		panic(err)
	}

	switch firstBytes[0] {
	case 'L':
		executeListCommand(in, figon)

	default:
		executeOperationWithList(in, figon)
	}
}

func executeListCommand(in *bufio.Reader, figon *Figon) {
	_, err := in.Discard(len("List "))
	if err != nil {
		panic(err)
	}

	data, err := in.ReadSlice(' ')
	if err != nil {
		panic(err)
	}

	newListName := [10]byte{}
	copy(newListName[:], data[:len(data)-1])

	_, err = in.Discard(len("= "))
	if err != nil {
		panic(err)
	}

	nextThreeBytes, err := in.Peek(4)
	if err != nil {
		panic(err)
	}

	switch [4]byte(nextThreeBytes) {
	case [4]byte{'n', 'e', 'w', ' '}:
		executeNewListCommand(in, newListName, figon)

	default:
		executeNewSubListCommand(in, newListName, figon)
	}
}

func executeNewListCommand(in *bufio.Reader, newListName [10]byte, figon *Figon) {
	_, err := in.Discard(len("new List("))
	if err != nil {
		panic(err)
	}

	figon.addNewList(CommandNewList{
		ListName: newListName,
		Digits:   scanDigits(in),
	})
}

func executeNewSubListCommand(in *bufio.Reader, newListName [10]byte, figon *Figon) {
	command := CommandNewSubList{}
	command.ListName = newListName

	data, err := in.ReadSlice('.')
	if err != nil {
		panic(err)
	}

	copy(command.RootListName[:], data[:len(data)-1])

	_, err = in.Discard(len("subList("))
	if err != nil {
		panic(err)
	}

	data, err = in.ReadSlice(',')
	if err != nil {
		panic(err)
	}

	data = data[:len(data)-1]
	command.From = fastAtoi(data)

	data, err = in.ReadSlice(')')
	if err != nil {
		panic(err)
	}

	data = data[:len(data)-1]

	command.To = fastAtoi(data)
	command.From--

	figon.addNewSubList(command)
}

func scanDigits(in *bufio.Reader) []int {
	digitsLine, err := in.ReadSlice(')')
	if err != nil {
		panic(err)
	}

	digitsLine = digitsLine[:len(digitsLine)-1]

	return fastAtoiSliceDelim(digitsLine)
}

func executeOperationWithList(in *bufio.Reader, figon *Figon) {
	data, err := in.ReadSlice('.')
	if err != nil {
		panic(err)
	}

	listName := [10]byte{}
	copy(listName[:], data[:len(data)-1])

	nextBytes, err := in.Peek(1)
	if err != nil {
		panic(err)
	}

	switch nextBytes[0] {
	case 's':
		executeSetCommand(in, listName, figon)

	case 'a':
		executeAddCommand(in, listName, figon)

	case 'g':
		executeGetCommand(in, listName, figon)

	default:
		panic("invalid symbol: " + string(nextBytes))
	}
}

func executeSetCommand(in *bufio.Reader, listName [10]byte, figon *Figon) {
	command := CommandSet{}
	command.ListName = listName

	_, err := in.Discard(len("set("))
	if err != nil {
		panic(err)
	}

	data, err := in.ReadSlice(',')
	if err != nil {
		panic(err)
	}

	data = data[:len(data)-1]
	command.Index = fastAtoi(data)

	data, err = in.ReadSlice(')')
	if err != nil {
		panic(err)
	}

	data = data[:len(data)-1]

	command.Value = fastAtoi(data)
	command.Index--

	figon.setInList(command)
}

func executeGetCommand(in *bufio.Reader, listName [10]byte, figon *Figon) {
	command := CommandGet{}
	command.ListName = listName

	_, err := in.Discard(len("get("))
	if err != nil {
		panic(err)
	}

	data, err := in.ReadSlice(')')
	if err != nil {
		panic(err)
	}

	data = data[:len(data)-1]

	command.Index = fastAtoi(data)
	command.Index--

	figon.writeListIndex(command)
}

func executeAddCommand(in *bufio.Reader, listName [10]byte, figon *Figon) {
	command := CommandAdd{}
	command.ListName = listName

	_, err := in.Discard(len("add("))
	if err != nil {
		panic(err)
	}

	data, err := in.ReadSlice(')')
	if err != nil {
		panic(err)
	}

	data = data[:len(data)-1]

	command.Value = fastAtoi(data)

	figon.addToList(command)
}

func fastAtoi(raw []byte) int {
	num := 0

	for _, sym := range raw {
		num = num*10 + int(sym-48)
	}

	return num
}

func fastAtoiSliceDelim(raw []byte) []int {
	digits := make([]int, 0)
	digit := 0

	for _, sym := range raw {
		if sym == ',' {
			digits = append(digits, digit)
			digit = 0

			continue
		}

		digit = digit*10 + int(sym-48)
	}

	digits = append(digits, digit)

	return digits
}
