// Generated by Fern. Do not edit.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

type UnionWithLiteral struct {
	typeName      string
	stringLiteral string
	String        string
}

func (u *UnionWithLiteral) StringLiteral() string {
	return u.stringLiteral
}

func (u *UnionWithLiteral) UnmarshalJSON(data []byte) error {
	var valueStringLiteral string
	if err := json.Unmarshal(data, &valueStringLiteral); err == nil {
		if valueStringLiteral == "fern" {
			u.typeName = "stringLiteral"
			u.stringLiteral = valueStringLiteral
			return nil
		}
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		u.typeName = "string"
		u.String = valueString
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, u)
}

func (u UnionWithLiteral) MarshalJSON() ([]byte, error) {
	switch u.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", u.typeName, u)
	case "stringLiteral":
		return json.Marshal("fern")
	case "string":
		return json.Marshal(u.String)
	}
}

type UnionWithLiteralVisitor interface {
	VisitStringLiteral(string) error
	VisitString(string) error
}

func (u *UnionWithLiteral) Accept(v UnionWithLiteralVisitor) error {
	switch u.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", u.typeName, u)
	case "stringLiteral":
		return v.VisitStringLiteral(u.stringLiteral)
	case "string":
		return v.VisitString(u.String)
	}
}