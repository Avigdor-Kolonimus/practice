package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Bank struct {
	accounts map[string]int
}

func NewBank() *Bank {
	return &Bank{
		accounts: map[string]int{},
	}
}

func (b *Bank) Deposit(name string, sum int) {
	b.accounts[name] += sum
}

func (b *Bank) Withdraw(name string, sum int) {
	b.accounts[name] -= sum
}

func (b *Bank) Balance(name string) (int, bool) {
	v, ok := b.accounts[name]
	return v, ok
}

func (b *Bank) Transfer(name1, name2 string, sum int) {
	b.accounts[name1] -= sum
	b.accounts[name2] += sum
}

func (b *Bank) Income(p int) {
	pr := 1 + float64(p)/100
	for k, v := range b.accounts {
		if v > 0 {
			b.accounts[k] = int(float64(v) * pr)
		}
	}
}

// https://coderun.yandex.ru/problem/bank-accounts
// BankAccounts - problem 72
func BankAccounts() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	bank := NewBank()
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimRight(line, "\r\n")
		tokens := strings.Fields(line)

		command := tokens[0]
		switch command {

		case "DEPOSIT":
			name := tokens[1]
			sum, err := strconv.Atoi(tokens[2])
			if err != nil {
				panic(err)
			}

			bank.Deposit(name, sum)

		case "WITHDRAW":
			name := tokens[1]
			sum, err := strconv.Atoi(tokens[2])
			if err != nil {
				panic(err)
			}

			bank.Withdraw(name, sum)

		case "BALANCE":
			name := tokens[1]
			balance, exists := bank.Balance(name)

			if !exists {
				writer.WriteString("ERROR")
			} else {
				writer.WriteString(strconv.Itoa(balance))
			}
			writer.WriteByte('\n')

		case "TRANSFER":
			from := tokens[1]
			to := tokens[2]
			sum, err := strconv.Atoi(tokens[3])
			if err != nil {
				panic(err)
			}

			bank.Transfer(from, to, sum)

		case "INCOME":
			p, err := strconv.Atoi(tokens[1])
			if err != nil {
				panic(err)
			}

			bank.Income(p)
		}
	}
}
