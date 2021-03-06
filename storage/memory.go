package storage

import (
	"github.com/xwb1989/shortener/storage/encoder"
)

type memMap struct {
	table   map[uint64]string
	encoder encoder.Encoder
}

func (m *memMap) Read(key string) (string, error) {
	k, err := m.encoder.StringToKey(key)
	if err != nil {
		return "", err
	}
	if val, ok := m.table[k]; ok {
		return val, nil
	} else {
		return val, InvalidKeyError(key)
	}
}

func (m *memMap) Write(url string) (string, error) {
	key := m.encoder.Encode(url)
	m.table[key] = url
	return m.encoder.KeyToString(key), nil
}

func NewMemMap(encoder encoder.Encoder) Storage {
	return &memMap{table: map[uint64]string{}, encoder: encoder}
}
