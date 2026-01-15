package main

import (
	"fmt"

	"example.com/test-project/internal/foo"
)

// This should trigger the linter
func ExposeInternal(f foo.InternalType) {
	fmt.Println(f.Val)
}

// This should NOT trigger (internal usage)
func useInternal(f foo.InternalType) {
	fmt.Println(f.Val)
}

func main() {
	useInternal(foo.InternalType{Val: 42})
}
