package testutil

import (
	"encoding/json"
	"sort"
)

// CanonicalJSON canonicalizes the JSON data such that if two JSON documents `a` and `b` are
// equivalent (ignoring object property ordering) then canonicalizeJSON(a) == canonicalizeJSON(b).
func CanonicalJSON(data []byte) []byte {
	var root interface{}
	if err := json.Unmarshal(data, &root); err != nil {
		panic(err)
	}

	visit := func(vp *interface{}) {
		switch v := (*vp).(type) {
		case map[string]interface{}:
			keys := make([]string, 0, len(v))
			for k := range v {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			m := make([]interface{}, 2*len(keys))
			for i, k := range keys {
				m[2*i] = k
				m[2*i+1] = v[k]
			}
			*vp = m
		}
	}
	visit(&root)
	data, err := json.Marshal(root)
	if err != nil {
		panic(err)
	}
	return data
}
