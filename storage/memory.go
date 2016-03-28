package storage

type memMap struct {
	table map[string]string
}

func (m *memMap) Read(key string) (string, error) {
	if val, ok := m.table[key]; ok {
		return val, nil
	} else {
		return val, InvalidKeyError()
	}
}

func (m *memMap) Write(key, val string) error {
	m.table[key] = val
	return nil
}

func NewMemMap() Storage {
	return &memMap{table: map[string]string{}}
}
