// Generated by Fern. Do not edit.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

type UnionWithUnknown struct {
	Type    string
	Foo     *Foo
	Unknown any
}

func (u *UnionWithUnknown) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	u.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "foo":
		value := new(Foo)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Foo = value
	case "unknown":
		value := make(map[string]any)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Unknown = value
	}
	return nil
}

func (u UnionWithUnknown) MarshalJSON() ([]byte, error) {
	switch u.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", u.Type, u)
	case "foo":
		var marshaler = struct {
			Type string `json:"type"`
			*Foo
		}{
			Type: u.Type,
			Foo:  u.Foo,
		}
		return json.Marshal(marshaler)
	case "unknown":
		var marshaler = struct {
			Type    string `json:"type"`
			Unknown any    `json:"unknown"`
		}{
			Type:    u.Type,
			Unknown: u.Unknown,
		}
		return json.Marshal(marshaler)
	}
}

type UnionWithUnknownVisitor interface {
	VisitFoo(*Foo) error
	VisitUnknown(any) error
}

func (u *UnionWithUnknown) Accept(v UnionWithUnknownVisitor) error {
	switch u.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", u.Type, u)
	case "foo":
		return v.VisitFoo(u.Foo)
	case "unknown":
		return v.VisitUnknown(u.Unknown)
	}
}