package storage

type memMap struct {
	table map[int]string
}

func (m *memMap) Read(key int) (string, error) {
	if val, ok := m.table[key]; ok {
		return val, nil
	} else {
		return val, InvalidKeyError()
	}
}

func (m *memMap) Write(key int, val string) error {
	m.table[key] = val
	return nil
}

func NewMemMap() Storage {
	return &memMap{table: map[int]string{}}
}
