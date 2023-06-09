// Generated by Fern. Do not edit.

package ir

import (
	json "encoding/json"
	fmt "fmt"
)

type TypeReference struct {
	Type      string
	Container *ContainerType
	Named     *DeclaredTypeName
	Primitive PrimitiveType
	Unknown   any
}

func (t *TypeReference) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"_type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	t.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "container":
		var valueUnmarshaler struct {
			Container *ContainerType `json:"container"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		t.Container = valueUnmarshaler.Container
	case "named":
		value := new(DeclaredTypeName)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Named = value
	case "primitive":
		var valueUnmarshaler struct {
			Primitive PrimitiveType `json:"primitive"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		t.Primitive = valueUnmarshaler.Primitive
	case "unknown":
		value := make(map[string]any)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Unknown = value
	}
	return nil
}

func (t TypeReference) MarshalJSON() ([]byte, error) {
	switch t.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", t.Type, t)
	case "container":
		var marshaler = struct {
			Type      string         `json:"_type"`
			Container *ContainerType `json:"container"`
		}{
			Type:      t.Type,
			Container: t.Container,
		}
		return json.Marshal(marshaler)
	case "named":
		var marshaler = struct {
			Type string `json:"_type"`
			*DeclaredTypeName
		}{
			Type:             t.Type,
			DeclaredTypeName: t.Named,
		}
		return json.Marshal(marshaler)
	case "primitive":
		var marshaler = struct {
			Type      string        `json:"_type"`
			Primitive PrimitiveType `json:"primitive"`
		}{
			Type:      t.Type,
			Primitive: t.Primitive,
		}
		return json.Marshal(marshaler)
	case "unknown":
		var marshaler = struct {
			Type    string `json:"_type"`
			Unknown any    `json:"unknown"`
		}{
			Type:    t.Type,
			Unknown: t.Unknown,
		}
		return json.Marshal(marshaler)
	}
}

type TypeReferenceVisitor interface {
	VisitContainer(*ContainerType) error
	VisitNamed(*DeclaredTypeName) error
	VisitPrimitive(PrimitiveType) error
	VisitUnknown(any) error
}

func (t *TypeReference) Accept(v TypeReferenceVisitor) error {
	switch t.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", t.Type, t)
	case "container":
		return v.VisitContainer(t.Container)
	case "named":
		return v.VisitNamed(t.Named)
	case "primitive":
		return v.VisitPrimitive(t.Primitive)
	case "unknown":
		return v.VisitUnknown(t.Unknown)
	}
}
