package p

import "encoding/json"

type ObjectWithProps struct {
	A          string         `json:"a,omitempty"`
	B          string         `json:"b,omitempty"`
	Additional map[string]any `json:"-"` // additionalProperties not explicitly defined in the schema
}

func (v ObjectWithProps) MarshalJSON() ([]byte, error) {
	m := make(map[string]any, len(v.Additional)+2)
	for k, v := range v.Additional {
		m[k] = v
	}
	m["a"] = v.A
	m["b"] = v.B
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
	delete(m, "a")
	delete(m, "b")
	if len(m) > 0 {
		v.Additional = make(map[string]any, len(m))
	}
	for k, vv := range m {
		v.Additional[k] = vv
	}
	return nil
}
