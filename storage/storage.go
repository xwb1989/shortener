package storage

type Reader interface {
	Read(string) (string, error)
}

type Writer interface {
	Write(string, string) error
}
