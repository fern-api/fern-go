// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type Environments uint8

const (
	EnvironmentsInside Environments = iota + 1
	EnvironmentsOutside
)

func (e Environments) String() string {
	switch e {
	default:
		return strconv.Itoa(int(e))
	case EnvironmentsInside:
		return "Inside"
	case EnvironmentsOutside:
		return "Outside"
	}
}

func (e Environments) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", e.String())), nil
}

func (e *Environments) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "Inside":
		value := EnvironmentsInside
		*e = value
	case "Outside":
		value := EnvironmentsOutside
		*e = value
	}
	return nil
}
