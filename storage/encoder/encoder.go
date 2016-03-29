package encoder

type Encoder interface {
	// Encode hash the input string to uint64 as key
	Encode(string) uint64

	// KeyToString converts int to its string representation
	KeyToString(uint64) string

	// StringToKey converts key string representation to int
	StringToKey(string) (uint64, error)
}
