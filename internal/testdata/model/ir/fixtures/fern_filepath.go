// This file was auto-generated by Fern from our API Definition.

package ir

type FernFilepath struct {
	AllParts    []*Name `json:"allParts,omitempty"`
	PackagePath []*Name `json:"packagePath,omitempty"`
	File        *Name   `json:"file,omitempty"`
}
