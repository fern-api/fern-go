// This file was auto-generated by Fern from our API Definition.

package ir

import (
	json "encoding/json"
	fmt "fmt"
)

type ErrorDiscriminationStrategy struct {
	Type       string
	StatusCode any
	Property   *ErrorDiscriminationByPropertyStrategy
}

func NewErrorDiscriminationStrategyFromStatusCode(value any) *ErrorDiscriminationStrategy {
	return &ErrorDiscriminationStrategy{Type: "statusCode", StatusCode: value}
}

func NewErrorDiscriminationStrategyFromProperty(value *ErrorDiscriminationByPropertyStrategy) *ErrorDiscriminationStrategy {
	return &ErrorDiscriminationStrategy{Type: "property", Property: value}
}

func (e *ErrorDiscriminationStrategy) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	e.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "statusCode":
		value := make(map[string]any)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.StatusCode = value
	case "property":
		value := new(ErrorDiscriminationByPropertyStrategy)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.Property = value
	}
	return nil
}

func (e ErrorDiscriminationStrategy) MarshalJSON() ([]byte, error) {
	switch e.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", e.Type, e)
	case "statusCode":
		var marshaler = struct {
			Type       string `json:"type"`
			StatusCode any    `json:"statusCode,omitempty"`
		}{
			Type:       e.Type,
			StatusCode: e.StatusCode,
		}
		return json.Marshal(marshaler)
	case "property":
		var marshaler = struct {
			Type string `json:"type"`
			*ErrorDiscriminationByPropertyStrategy
		}{
			Type:                                  e.Type,
			ErrorDiscriminationByPropertyStrategy: e.Property,
		}
		return json.Marshal(marshaler)
	}
}

type ErrorDiscriminationStrategyVisitor interface {
	VisitStatusCode(any) error
	VisitProperty(*ErrorDiscriminationByPropertyStrategy) error
}

func (e *ErrorDiscriminationStrategy) Accept(v ErrorDiscriminationStrategyVisitor) error {
	switch e.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", e.Type, e)
	case "statusCode":
		return v.VisitStatusCode(e.StatusCode)
	case "property":
		return v.VisitProperty(e.Property)
	}
}
