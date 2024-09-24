package dsa

type KeyToValuesStore struct {
	key    string
	values []string
}

func NewKeyToValuesStore(key string, values ...string) *KeyToValuesStore {
	return &KeyToValuesStore{key: key, values: values}
}

func (kv *KeyToValuesStore) Key() string {
	return kv.key
}

func (kv *KeyToValuesStore) Values() []string {
	return kv.values
}

func (kv *KeyToValuesStore) ToMap() map[string][]string {
	m := make(map[string][]string)

	m[kv.Key()] = kv.Values()

	return m
}
