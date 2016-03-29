package encoder

import (
	"strconv"
)

type incrementer struct {
	id uint64
}

func (i *incrementer) Encode(s string) uint64 {
	ret := i.id
	i.id++
	return ret
}

func (i *incrementer) KeyToString(key uint64) string {
	return strconv.FormatUint(key, 36)
}

func (i *incrementer) StringToKey(key string) (uint64, error) {
	return strconv.ParseUint(key, 36, 64)
}

func NewIncrementalEncoder(start uint64) Encoder {
	return &incrementer{id: start}
}
