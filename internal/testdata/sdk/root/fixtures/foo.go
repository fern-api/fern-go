// This file was auto-generated by Fern from our API Definition.

package api

type Foo struct {
	Id      *Id      `json:"id,omitempty"`
	Name    *string  `json:"name,omitempty"`
	List    *string  `json:"list,omitempty"`
	Type    *FooType `json:"type,omitempty"`
	Request *Request `json:"request,omitempty"`
	Delay   *string  `json:"delay,omitempty"`
}
