package p

type NullableArrayElement struct {
	A string  `json:"a,omitempty"`
	B float64 `json:"b,omitempty"`
}
type Wrapper struct {
	Array []*NullableArrayElement `json:"array,omitempty"`
}
