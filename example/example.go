package example

//go:generate go-enumerator
// Kind demonstrates integer style enums
type Kind int

const (
	Kind1 Kind = iota
	Kind2
)

//go:generate go-enumerator
// StrKind demonstrates string style enums
type StrKind string

const (
	Hello StrKind = "Hello"
	World StrKind = "World"
)
