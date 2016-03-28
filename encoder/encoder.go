package encoder

type Encoder interface {
	// Encode converts string into base64 number
	Encode(string) int

	// ToString converts int to its string representation
	ToString(int) string

	// FromString converts encoded string representation to int
	FromString(string) int
}
