package encoder

type Encoder interface {
	// Encode hash the input string to an integer as key
	Encode(string) int

	// KeyToString converts int to its string representation
	KeyToString(int) string

	// StringToKey converts key string representation to int
	StringToKey(string) int
}
