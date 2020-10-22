package config

import (
	"otter/pkg/jwt"

	"github.com/EricChiou/config"
)

// Config struct, set parameter in config.ini file
type Config struct {
	ServerName      string `key:"SERVER_NAME"`
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
	ConfigPath string     = "./config.ini"
	JwtAlg     jwt.AlgTyp = jwt.HS256
	Sha3Len    int        = 256
)

var cfg = Config{}

// Load config from config.ini
func Load(filePath string) error {
	return config.Load(filePath, &cfg)
}

// Get config
func Get() Config {
	return cfg
}
