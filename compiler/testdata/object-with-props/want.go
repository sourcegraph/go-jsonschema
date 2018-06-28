package p

import "encoding/json"

type ObjectWithProps struct {
	A          string                 `json:"a,omitempty"`
	Additional map[string]interface{} `json:"-"`
}

func (v ObjectWithProps) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, len(v.Additional)+1)
	for k, v := range v.Additional {
		m[k] = v
	}
	m["a"] = v.A
	return json.Marshal(m)
}
func (v *ObjectWithProps) UnmarshalJSON(data []byte) error {
	var s struct {
		A string `json:"a,omitempty"`
	}
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*v = ObjectWithProps{A: s.A}
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	delete(m, "a")
	if len(m) > 0 {
		(*v).Additional = make(map[string]interface{}, len(m))
	}
	for k, vv := range m {
		(*v).Additional[k] = vv
	}
	return nil
}
