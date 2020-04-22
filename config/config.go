package config

import (
	conf "otter/pkg/config"
	"reflect"
)

// Config config struct
type config struct {
	ServerPort      string `key:"SERVER_PORT"`
	SSLCertFilePath string `key:"SSL_CERT_FILE_PATH"`
	SSLKeyFilePath  string `key:"SSL_KEY_FILE_PATH"`
	MySQLAddr       string `key:"MYSQL_ADDR"`
	MySQLPort       string `key:"MYSQL_PORT"`
	MySQLUserName   string `key:"MYSQL_USERNAME"`
	MySQLPassword   string `key:"MYSQL_PASSWORD"`
	MySQLDBNAME     string `key:"MYSQL_DBNAME"`
	JWTKey          string `key:"JWT_KEY"`
	JWTExpire       string `key:"JWT_EXPIRE"`
	ENV             string `key:"ENV"`
}

var cfg = config{}

// LoadConfig load config
func Load(configFilePath string) error {
	keyValue, err := conf.LoadConfigFile(configFilePath)
	if err != nil {
		return err
	}

	keys := reflect.TypeOf(cfg)
	values := reflect.ValueOf(&cfg).Elem()
	for i := 0; i < keys.NumField(); i++ {
		values.Field(i).SetString(keyValue[keys.Field(i).Tag.Get("key")])
	}
	return nil
}

// Get config
func Get() config {
	return cfg
}
