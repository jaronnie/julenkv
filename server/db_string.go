package server

func (db *JulenKv) Set(key, value []byte) error {
	if err := db.dictIdx.Set(key, value); err != nil {
		return err
	}
	return nil
}

func (db *JulenKv) Get(key []byte) ([]byte, error) {
	return db.dictIdx.Get(key), nil
}
