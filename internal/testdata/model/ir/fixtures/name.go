// This file was auto-generated by Fern from our API Definition.

package ir

type Name struct {
	OriginalName       string               `json:"originalName"`
	CamelCase          *SafeAndUnsafeString `json:"camelCase,omitempty"`
	PascalCase         *SafeAndUnsafeString `json:"pascalCase,omitempty"`
	SnakeCase          *SafeAndUnsafeString `json:"snakeCase,omitempty"`
	ScreamingSnakeCase *SafeAndUnsafeString `json:"screamingSnakeCase,omitempty"`
}
