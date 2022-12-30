package p

import "encoding/json"

type ObjectWithProps struct {
	P0         string         `json:"p0,omitempty"`
	P1         float64        `json:"p1"`
	P2         *P2            `json:"p2,omitempty"`
	Additional map[string]any `json:"-"` // additionalProperties not explicitly defined in the schema
}

func (v ObjectWithProps) MarshalJSON() ([]byte, error) {
	m := make(map[string]any, len(v.Additional))
	for k, v := range v.Additional {
		m[k] = v
	}
	type wrapper ObjectWithProps
	b, err := json.Marshal(wrapper(v))
	if err != nil {
		return nil, err
	}
	var m2 map[string]any
	if err := json.Unmarshal(b, &m2); err != nil {
		return nil, err
	}
	for k, v := range m2 {
		m[k] = v
	}
	return json.Marshal(m)
}
func (v *ObjectWithProps) UnmarshalJSON(data []byte) error {
	type wrapper ObjectWithProps
	var s wrapper
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*v = ObjectWithProps(s)
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	delete(m, "p0")
	delete(m, "p1")
	delete(m, "p2")
	if len(m) > 0 {
		v.Additional = make(map[string]any, len(m))
	}
	for k, vv := range m {
		v.Additional[k] = vv
	}
	return nil
}

type P2 struct {
	A string `json:"a,omitempty"`
}
