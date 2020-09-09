package p

type A struct {
	// B description: d
	B string `json:"b,omitempty"`
}
type RefToPrimitive struct {
	A string `json:"a,omitempty"`
	B string `json:"b,omitempty"`
}
