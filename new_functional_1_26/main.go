package main

import (
	"fmt"

	exp "new_functional_1_26/examples"
)

func main() {
	fmt.Println("new()")
	exp.PtrNew()

	fmt.Println("errors")
	exp.ErrType()

	fmt.Println("hpke")
	exp.Exphpke()
	exp.ExphpkeSenderRecipient()
	exp.ExphpkeKEM()

	fmt.Println("buf.peek")
	exp.PeekBuffer()

	fmt.Println("signal.NotifyContext")
	exp.NotifyContext()
}
