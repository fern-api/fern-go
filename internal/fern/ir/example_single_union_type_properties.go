// Generated by Fern. Do not edit.

package ir

import (
	json "encoding/json"
	fmt "fmt"
)

type ExampleSingleUnionTypeProperties struct {
	Type                   string
	SamePropertiesAsObject *ExampleNamedType
	SingleProperty         *ExampleTypeReference
	NoProperties           any
}

func (e *ExampleSingleUnionTypeProperties) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	e.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "samePropertiesAsObject":
		value := new(ExampleNamedType)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.SamePropertiesAsObject = value
	case "singleProperty":
		value := new(ExampleTypeReference)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.SingleProperty = value
	case "noProperties":
		value := make(map[string]any)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.NoProperties = value
	}
	return nil
}

func (e ExampleSingleUnionTypeProperties) MarshalJSON() ([]byte, error) {
	switch e.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", e.Type, e)
	case "samePropertiesAsObject":
		var marshaler = struct {
			Type string `json:"type"`
			*ExampleNamedType
		}{
			Type:             e.Type,
			ExampleNamedType: e.SamePropertiesAsObject,
		}
		return json.Marshal(marshaler)
	case "singleProperty":
		var marshaler = struct {
			Type string `json:"type"`
			*ExampleTypeReference
		}{
			Type:                 e.Type,
			ExampleTypeReference: e.SingleProperty,
		}
		return json.Marshal(marshaler)
	case "noProperties":
		var marshaler = struct {
			Type         string `json:"type"`
			NoProperties any    `json:"noProperties"`
		}{
			Type:         e.Type,
			NoProperties: e.NoProperties,
		}
		return json.Marshal(marshaler)
	}
}

type ExampleSingleUnionTypePropertiesVisitor interface {
	VisitSamePropertiesAsObject(*ExampleNamedType) error
	VisitSingleProperty(*ExampleTypeReference) error
	VisitNoProperties(any) error
}

func (e *ExampleSingleUnionTypeProperties) Accept(v ExampleSingleUnionTypePropertiesVisitor) error {
	switch e.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", e.Type, e)
	case "samePropertiesAsObject":
		return v.VisitSamePropertiesAsObject(e.SamePropertiesAsObject)
	case "singleProperty":
		return v.VisitSingleProperty(e.SingleProperty)
	case "noProperties":
		return v.VisitNoProperties(e.NoProperties)
	}
}
