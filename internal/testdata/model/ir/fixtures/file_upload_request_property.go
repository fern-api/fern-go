// Generated by Fern. Do not edit.

package ir

import (
	json "encoding/json"
	fmt "fmt"
)

type FileUploadRequestProperty struct {
	Type         string
	File         *FileProperty
	BodyProperty *InlinedRequestBodyProperty
}

func (f *FileUploadRequestProperty) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	f.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "file":
		value := new(FileProperty)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		f.File = value
	case "bodyProperty":
		value := new(InlinedRequestBodyProperty)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		f.BodyProperty = value
	}
	return nil
}

func (f FileUploadRequestProperty) MarshalJSON() ([]byte, error) {
	switch f.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", f.Type, f)
	case "file":
		var marshaler = struct {
			Type string `json:"type"`
			*FileProperty
		}{
			Type:         f.Type,
			FileProperty: f.File,
		}
		return json.Marshal(marshaler)
	case "bodyProperty":
		var marshaler = struct {
			Type string `json:"type"`
			*InlinedRequestBodyProperty
		}{
			Type:                       f.Type,
			InlinedRequestBodyProperty: f.BodyProperty,
		}
		return json.Marshal(marshaler)
	}
}

type FileUploadRequestPropertyVisitor interface {
	VisitFile(*FileProperty) error
	VisitBodyProperty(*InlinedRequestBodyProperty) error
}

func (f *FileUploadRequestProperty) Accept(v FileUploadRequestPropertyVisitor) error {
	switch f.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", f.Type, f)
	case "file":
		return v.VisitFile(f.File)
	case "bodyProperty":
		return v.VisitBodyProperty(f.BodyProperty)
	}
}