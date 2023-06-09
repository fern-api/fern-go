// This file was auto-generated by Fern from our API Definition.

package ir

import (
	json "encoding/json"
	fmt "fmt"
)

type Environments struct {
	Type             string
	SingleBaseUrl    *SingleBaseUrlEnvironments
	MultipleBaseUrls *MultipleBaseUrlsEnvironments
}

func NewEnvironmentsFromSingleBaseUrl(value *SingleBaseUrlEnvironments) *Environments {
	return &Environments{Type: "singleBaseUrl", SingleBaseUrl: value}
}

func NewEnvironmentsFromMultipleBaseUrls(value *MultipleBaseUrlsEnvironments) *Environments {
	return &Environments{Type: "multipleBaseUrls", MultipleBaseUrls: value}
}

func (e *Environments) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	e.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "singleBaseUrl":
		value := new(SingleBaseUrlEnvironments)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.SingleBaseUrl = value
	case "multipleBaseUrls":
		value := new(MultipleBaseUrlsEnvironments)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.MultipleBaseUrls = value
	}
	return nil
}

func (e Environments) MarshalJSON() ([]byte, error) {
	switch e.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", e.Type, e)
	case "singleBaseUrl":
		var marshaler = struct {
			Type string `json:"type"`
			*SingleBaseUrlEnvironments
		}{
			Type:                      e.Type,
			SingleBaseUrlEnvironments: e.SingleBaseUrl,
		}
		return json.Marshal(marshaler)
	case "multipleBaseUrls":
		var marshaler = struct {
			Type string `json:"type"`
			*MultipleBaseUrlsEnvironments
		}{
			Type:                         e.Type,
			MultipleBaseUrlsEnvironments: e.MultipleBaseUrls,
		}
		return json.Marshal(marshaler)
	}
}

type EnvironmentsVisitor interface {
	VisitSingleBaseUrl(*SingleBaseUrlEnvironments) error
	VisitMultipleBaseUrls(*MultipleBaseUrlsEnvironments) error
}

func (e *Environments) Accept(v EnvironmentsVisitor) error {
	switch e.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", e.Type, e)
	case "singleBaseUrl":
		return v.VisitSingleBaseUrl(e.SingleBaseUrl)
	case "multipleBaseUrls":
		return v.VisitMultipleBaseUrls(e.MultipleBaseUrls)
	}
}
