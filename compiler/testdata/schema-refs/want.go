package p

type D struct {
	E      *E   `json:"e,omitempty"`
	EArray []*E `json:"eArray"`
}
type E struct {
	A string `json:"a,omitempty"`
}
type SchemaRefs struct {
	D             D    `json:"d"`
	DArray        []*D `json:"dArray"`
	DArrayPointer []*D `json:"dArrayPointer,omitempty"`
	DPointer      *D   `json:"dPointer,omitempty"`
}
