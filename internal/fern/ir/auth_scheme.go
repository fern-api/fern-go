// Generated by Fern. Do not edit.

package ir

import (
	json "encoding/json"
	fmt "fmt"
)

type AuthScheme struct {
	Type   string
	Bearer *BearerAuthScheme
	Basic  *BasicAuthScheme
	Header *HeaderAuthScheme
}

func (a *AuthScheme) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"_type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	a.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "bearer":
		value := new(BearerAuthScheme)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Bearer = value
	case "basic":
		value := new(BasicAuthScheme)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Basic = value
	case "header":
		value := new(HeaderAuthScheme)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Header = value
	}
	return nil
}

func (a AuthScheme) MarshalJSON() ([]byte, error) {
	switch a.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", a.Type, a)
	case "bearer":
		var marshaler = struct {
			Type string `json:"_type"`
			*BearerAuthScheme
		}{
			Type:             a.Type,
			BearerAuthScheme: a.Bearer,
		}
		return json.Marshal(marshaler)
	case "basic":
		var marshaler = struct {
			Type string `json:"_type"`
			*BasicAuthScheme
		}{
			Type:            a.Type,
			BasicAuthScheme: a.Basic,
		}
		return json.Marshal(marshaler)
	case "header":
		var marshaler = struct {
			Type string `json:"_type"`
			*HeaderAuthScheme
		}{
			Type:             a.Type,
			HeaderAuthScheme: a.Header,
		}
		return json.Marshal(marshaler)
	}
}

type AuthSchemeVisitor interface {
	VisitBearer(*BearerAuthScheme) error
	VisitBasic(*BasicAuthScheme) error
	VisitHeader(*HeaderAuthScheme) error
}

func (a *AuthScheme) Accept(v AuthSchemeVisitor) error {
	switch a.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", a.Type, a)
	case "bearer":
		return v.VisitBearer(a.Bearer)
	case "basic":
		return v.VisitBasic(a.Basic)
	case "header":
		return v.VisitHeader(a.Header)
	}
}
