package encoder

type Encoder interface {
	Encode(string) string
	Decode(string) string
}
