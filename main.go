// go-enumerator is a tool designed to be called by go:generate for generating enum-like
// code from constants.
//
// Be default, go-enumerator will look for a type definition immediately following
// the go:generate statement from which it was called, and will generate methods for
// that type.
//
// For example, given code similar to what is shown below
//
// 		//go:generate go-enumerator
// 		type Kind int
//
// 		const (
//			Kind1
//			Kind2
// 		)
//
// go-enumerator will generate implementations for the following methods
//
// 		// String implements fmt.Stringer
// 		func (k Kind) String() string { /* omitted for brevity */ }
//
// 		// Scan implements fmt.Scanner
// 		func (k *Kind) Scan(ss fmt.ScanState, verb rune) error { /* omitted for brevity */ }
//
// 		// Defined returns true if k holds a defined value
// 		func (k Kind) Defined() bool { /* omitted for brevity */ }
//
// 		// Next returns the next defined value after k
// 		func (k Kind) Next() Kind { /* omitted for brevity */ }
//
// Hopefully, the default behavior will serve your needs, but if not it can be changed
// by supplying command-line arguments to the program. For help with the cli, run with the
// --help argument.
//
// 		go-enumerator --help
//
// Enjoy 😀
//
package main

import (
	"github.com/ajjensen13/go-enumerator/internal/cmd"
)

func main() {
	cmd.Execute()
}
