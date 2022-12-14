// Code generated by "go-enumerator"; DO NOT EDIT.

package example

import "fmt"

// String implements fmt.Stringer. If !k.Defined(), then a generated string is returned based on k's value.
func (k Kind) String() string {
	switch k {
	case Kind1:
		return "Kind1"
	case Kind2:
		return "Kind2"
	}
	return fmt.Sprintf("Kind(%d)", k)
}

// Bytes returns a byte-level representation of String(). If !k.Defined(), then a generated string is returned based on k's value.
func (k Kind) Bytes() []byte {
	switch k {
	case Kind1:
		return []byte{'K', 'i', 'n', 'd', '1'}
	case Kind2:
		return []byte{'K', 'i', 'n', 'd', '2'}
	}
	return []byte(fmt.Sprintf("Kind(%d)", k))
}

// Defined returns true if k holds a defined value.
func (k Kind) Defined() bool {
	switch k {
	case 0, 1:
		return true
	default:
		return false
	}
}

// Scan implements fmt.Scanner. Use fmt.Scan() to parse strings into Kind values
func (k *Kind) Scan(scanState fmt.ScanState, verb rune) error {
	token, err := scanState.Token(true, nil)
	if err != nil {
		return err
	}

	switch string(token) {
	case "Kind1":
		*k = Kind1
	case "Kind2":
		*k = Kind2
	default:
		return fmt.Errorf("unknown Kind value: %s", token)
	}
	return nil
}

// Next returns the next defined Kind. If k is not defined, then Next returns the first defined value.
// Next() can be used to loop through all values of an enum.
//
//	k := Kind(0)
//	for {
//		fmt.Println(k)
//		k = k.Next()
//		if k == Kind(0) {
//			break
//		}
//	}
//
// The exact order that values are returned when looping should not be relied upon.
func (k Kind) Next() Kind {
	switch k {
	case Kind1:
		return Kind2
	case Kind2:
		return Kind1
	default:
		return Kind1
	}
}

func _() {
	var x [1]struct{}
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the go-enumerator command to generate them again.
	_ = x[Kind1-0]
	_ = x[Kind2-1]
}

// MarshalJSON implements json.Marshaler
func (k Kind) MarshalJSON() ([]byte, error) {
	x := k.Bytes()
	y := make([]byte, 0, len(x))
	return append(append(append(y, '"'), x...), '"'), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (k *Kind) UnmarshalJSON(x []byte) error {
	switch string(x) {
	case "\"Kind1\"":
		*k = Kind1
		return nil
	case "\"Kind2\"":
		*k = Kind2
		return nil
	default:
		return fmt.Errorf("failed to parse value %v into %T", x, *k)
	}
}
