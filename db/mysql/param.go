package mysql

type sqlParams struct {
	kv map[string]interface{}
}

// Add value with key
func (params *sqlParams) Add(key string, value interface{}) {
	params.kv[key] = value
}

// Get value with key
func (params *sqlParams) Get(key string) interface{} {
	return params.kv[key]
}

// Remove value which key equal to input parameter
func (params *sqlParams) Remove(key string) {
	delete(params.kv, key)
}

// GetSQLParamsInstance get sqlParams instance
func GetSQLParamsInstance() sqlParams {
	return sqlParams{kv: map[string]interface{}{}}
}
