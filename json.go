package swiss

import (
	"errors"

	json "github.com/bytedance/sonic"
)

// MarshalJSON 自定义序列化
func (m *Map[K, V]) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte(`{}`), nil
	}
	m1 := make(map[K]V, m.Count())
	m.Iter(func(k K, v V) (stop bool) {
		m1[k] = v
		return
	})

	return json.Marshal(m1)
}

// UnmarshalJSON 自定义反序列化
func (m *Map[K, V]) UnmarshalJSON(data []byte) error {
	m1 := make(map[K]V, m.Count())
	if err := json.Unmarshal(data, m1); err != nil {
		return err
	}

	if m == nil {
		return errors.New("Map is nil, should be init first")
	}
	for k, v := range m1 {
		m.Put(k, v)
	}
	return nil
}
