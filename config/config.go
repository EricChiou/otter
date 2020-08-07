package config

import (
	cons "otter/constants"
	conf "otter/pkg/config"
	"reflect"
)

// config struct, set parameter in config.ini file
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

// config setting, set parameter here straightly
const (
	ConfigFilePath string         = "./config.ini"
	ServerName     string         = "otter framework"
	JwtAlg         cons.JwtAlgTyp = cons.JWTHS256
	Sha3Len        int            = 256
)

var cfg = config{}

// LoadConfig load config from config.ini
func Load(configFilePath string) error {
	keyValue, err := conf.LoadConfigFile(configFilePath)
	if err != nil {
		return err
	}

	keys := reflect.TypeOf(&cfg).Elem()
	values := reflect.ValueOf(&cfg).Elem()
	for i := 0; i < keys.NumField(); i++ {
		if len(keyValue[keys.Field(i).Tag.Get("key")]) != 0 {
			values.Field(i).SetString(keyValue[keys.Field(i).Tag.Get("key")])
		}
	}
	return nil
}

// Get config
func Get() config {
	return cfg
}
