// This file was auto-generated by Fern from our API Definition.

package ir

// A type, which is a name and a shape
type TypeDeclaration struct {
	Docs         *string           `json:"docs,omitempty"`
	Availability *Availability     `json:"availability,omitempty"`
	Name         *DeclaredTypeName `json:"name,omitempty"`
	Shape        *Type             `json:"shape,omitempty"`
	Examples     []*ExampleType    `json:"examples,omitempty"`
	// All other named types that this type references (directly or indirectly)
	ReferencedTypes []*DeclaredTypeName `json:"referencedTypes,omitempty"`
}
