package config

import (
	conf "otter/pkg/config"
	"reflect"
)

// ConfigKey config data keys, ConfigKey values need to match the config.ini keys.
var ConfigKey = config{
	ServerPort:      "SERVER_PORT",
	SSLCertFilePath: "SSL_CERT_FILE_PATH",
	SSLKeyFilePath:  "SSL_KEY_FILE_PATH",
	MySQLAddr:       "MYSQL_ADDR",
	MySQLPort:       "MYSQL_PORT",
	MySQLUserName:   "MYSQL_USERNAME",
	MySQLPassword:   "MYSQL_PASSWORD",
	MySQLDBNAME:     "MYSQL_DBNAME",
	JWTKey:          "JWT_KEY",
	JWTExpire:       "JWT_EXPIRE",
	ENV:             "ENV",
}

// Config config data
var Config = config{}

// LoadConfig load config
func LoadConfig(configFilePath string) error {
	keyValue, err := conf.LoadConfigFile(configFilePath)
	if err != nil {
		return err
	}

	keys := reflect.TypeOf(ConfigKey)
	values := reflect.ValueOf(ConfigKey)
	for i := 0; i < keys.NumField(); i++ {
		Config.setCfg(keys.Field(i).Name, keyValue[values.Field(i).Interface().(string)])
	}
	return nil
}
