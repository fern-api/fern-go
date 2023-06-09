// This file was auto-generated by Fern from our API Definition.

package ir

type UnionTypeDeclaration struct {
	Discriminant *NameAndWireValue `json:"discriminant,omitempty"`
	// A list of other types to inherit from
	Extends        []*DeclaredTypeName `json:"extends,omitempty"`
	Types          []*SingleUnionType  `json:"types,omitempty"`
	BaseProperties []*ObjectProperty   `json:"baseProperties,omitempty"`
}
