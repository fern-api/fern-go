// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type Something uint8

const (
	Somethingone Something = iota + 1
	SomethingOne
	SomethingONe
	SomethingONE
)

func (s Something) String() string {
	switch s {
	default:
		return strconv.Itoa(int(s))
	case Somethingone:
		return "one"
	case SomethingOne:
		return "One"
	case SomethingONe:
		return "ONe"
	case SomethingONE:
		return "ONE"
	}
}

func (s Something) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", s.String())), nil
}

func (s *Something) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "one":
		value := Somethingone
		*s = value
	case "One":
		value := SomethingOne
		*s = value
	case "ONe":
		value := SomethingONe
		*s = value
	case "ONE":
		value := SomethingONE
		*s = value
	}
	return nil
}
