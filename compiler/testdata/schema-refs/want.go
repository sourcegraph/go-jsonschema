package p

// D
type D struct {
	E      *E   `json:"e,omitempty"`
	EArray []*E `json:"eArray"`
}

// E
type E struct {
	A string `json:"a,omitempty"`
}

// SchemaRefs
type SchemaRefs struct {
	D             D    `json:"d"`
	DArray        []*D `json:"dArray"`
	DArrayPointer []*D `json:"dArrayPointer,omitempty"`
	DPointer      *D   `json:"dPointer,omitempty"`
}
