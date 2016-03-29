package encoder

type Encoder interface {
	// Encode hash the input string to an integer as key
	Encode(string) int

	// ToString converts int to its string representation
	KeyToString(int) string

	// FromString converts encoded string representation to int
	StringToKey(string) int
}
