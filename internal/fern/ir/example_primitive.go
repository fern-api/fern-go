// Generated by Fern. Do not edit.

package ir

import (
	json "encoding/json"
	fmt "fmt"
	uuid "github.com/gofrs/uuid/v5"
	time "time"
)

type ExamplePrimitive struct {
	Type     string
	Integer  int
	Double   float64
	String   string
	Boolean  bool
	Long     int64
	Datetime time.Time
	Date     time.Time
	Uuid     uuid.UUID
}

func (e *ExamplePrimitive) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	e.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "integer":
		var valueUnmarshaler struct {
			Integer int `json:"integer"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		e.Integer = valueUnmarshaler.Integer
	case "double":
		var valueUnmarshaler struct {
			Double float64 `json:"double"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		e.Double = valueUnmarshaler.Double
	case "string":
		var valueUnmarshaler struct {
			String string `json:"string"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		e.String = valueUnmarshaler.String
	case "boolean":
		var valueUnmarshaler struct {
			Boolean bool `json:"boolean"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		e.Boolean = valueUnmarshaler.Boolean
	case "long":
		var valueUnmarshaler struct {
			Long int64 `json:"long"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		e.Long = valueUnmarshaler.Long
	case "datetime":
		var valueUnmarshaler struct {
			Datetime time.Time `json:"datetime"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		e.Datetime = valueUnmarshaler.Datetime
	case "date":
		var valueUnmarshaler struct {
			Date time.Time `json:"date"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		e.Date = valueUnmarshaler.Date
	case "uuid":
		var valueUnmarshaler struct {
			Uuid uuid.UUID `json:"uuid"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		e.Uuid = valueUnmarshaler.Uuid
	}
	return nil
}

func (e ExamplePrimitive) MarshalJSON() ([]byte, error) {
	switch e.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", e.Type, e)
	case "integer":
		var marshaler = struct {
			Type    string `json:"type"`
			Integer int    `json:"integer"`
		}{
			Type:    e.Type,
			Integer: e.Integer,
		}
		return json.Marshal(marshaler)
	case "double":
		var marshaler = struct {
			Type   string  `json:"type"`
			Double float64 `json:"double"`
		}{
			Type:   e.Type,
			Double: e.Double,
		}
		return json.Marshal(marshaler)
	case "string":
		var marshaler = struct {
			Type   string `json:"type"`
			String string `json:"string"`
		}{
			Type:   e.Type,
			String: e.String,
		}
		return json.Marshal(marshaler)
	case "boolean":
		var marshaler = struct {
			Type    string `json:"type"`
			Boolean bool   `json:"boolean"`
		}{
			Type:    e.Type,
			Boolean: e.Boolean,
		}
		return json.Marshal(marshaler)
	case "long":
		var marshaler = struct {
			Type string `json:"type"`
			Long int64  `json:"long"`
		}{
			Type: e.Type,
			Long: e.Long,
		}
		return json.Marshal(marshaler)
	case "datetime":
		var marshaler = struct {
			Type     string    `json:"type"`
			Datetime time.Time `json:"datetime"`
		}{
			Type:     e.Type,
			Datetime: e.Datetime,
		}
		return json.Marshal(marshaler)
	case "date":
		var marshaler = struct {
			Type string    `json:"type"`
			Date time.Time `json:"date"`
		}{
			Type: e.Type,
			Date: e.Date,
		}
		return json.Marshal(marshaler)
	case "uuid":
		var marshaler = struct {
			Type string    `json:"type"`
			Uuid uuid.UUID `json:"uuid"`
		}{
			Type: e.Type,
			Uuid: e.Uuid,
		}
		return json.Marshal(marshaler)
	}
}

type ExamplePrimitiveVisitor interface {
	VisitInteger(int) error
	VisitDouble(float64) error
	VisitString(string) error
	VisitBoolean(bool) error
	VisitLong(int64) error
	VisitDatetime(time.Time) error
	VisitDate(time.Time) error
	VisitUuid(uuid.UUID) error
}

func (e *ExamplePrimitive) Accept(v ExamplePrimitiveVisitor) error {
	switch e.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", e.Type, e)
	case "integer":
		return v.VisitInteger(e.Integer)
	case "double":
		return v.VisitDouble(e.Double)
	case "string":
		return v.VisitString(e.String)
	case "boolean":
		return v.VisitBoolean(e.Boolean)
	case "long":
		return v.VisitLong(e.Long)
	case "datetime":
		return v.VisitDatetime(e.Datetime)
	case "date":
		return v.VisitDate(e.Date)
	case "uuid":
		return v.VisitUuid(e.Uuid)
	}
}
