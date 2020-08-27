package mysql

type sqlParams struct {
	kv map[string]string
}

// Add value with key
func (params *sqlParams) Add(key string, value string) {
	params.kv[key] = value
}

// Get value with key
func (params *sqlParams) Get(key string) string {
	return params.kv[key]
}

// Remove value which key equal to input parameter
func (params *sqlParams) Remove(key string) {
	delete(params.kv, key)
}

// GetSQLParamsInstance get sqlParams instance
func SQLParamsInstance() sqlParams {
	return sqlParams{kv: map[string]string{}}
}
