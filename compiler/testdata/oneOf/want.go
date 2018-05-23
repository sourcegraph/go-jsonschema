package p

import (
	"encoding/json"
	"errors"
	"fmt"
)

type A struct {
	B *B
	C *C
	D *D
}

func (v A) MarshalJSON() ([]byte, error) {
	if v.B != nil {
		return json.Marshal(v.B)
	}
	if v.C != nil {
		return json.Marshal(v.C)
	}
	if v.D != nil {
		return json.Marshal(v.D)
	}
	return nil, errors.New("tagged union type must have exactly 1 non-nil field value")
}
func (v *A) UnmarshalJSON(data []byte) error {
	var d struct {
		DiscriminantProperty string `json:"type"`
	}
	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}
	switch d.DiscriminantProperty {
	case "b":
		return json.Unmarshal(data, &v.B)
	case "c":
		return json.Unmarshal(data, &v.C)
	case "d":
		return json.Unmarshal(data, &v.D)
	}
	return fmt.Errorf("tagged union type must have a %q property whose value is one of %s", "type", []string{"b", "c", "d"})
}

type B struct {
	B    string `json:"b,omitempty"`
	Type string `json:"type"`
}
type C struct {
	C    bool   `json:"c,omitempty"`
	Type string `json:"type"`
}
type D struct {
	D    float64 `json:"d,omitempty"`
	Type string  `json:"type"`
}

// OneOf description: oneOf to implement a tagged union type
type OneOf struct {
	A []A `json:"a,omitempty"`
}
