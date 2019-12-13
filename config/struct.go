package config

// Config config struct
type config struct {
	ServerPort, SSLCertFilePath, SSLKeyFilePath                     string
	MySQLAddr, MySQLPort, MySQLUserName, MySQLPassword, MySQLDBNAME string // MySQL setting
	JWTKey, JWTExpire                                               string // JWT setting
	ENV                                                             string // Environment
}

func (cfg *config) setCfg(key, value string) {
	switch key {
	case "ServerPort":
		cfg.ServerPort = value
	case "SSLCertFilePath":
		cfg.SSLCertFilePath = value
	case "SSLKeyFilePath":
		cfg.SSLKeyFilePath = value
	case "MySQLAddr":
		cfg.MySQLAddr = value
	case "MySQLPort":
		cfg.MySQLPort = value
	case "MySQLUserName":
		cfg.MySQLUserName = value
	case "MySQLPassword":
		cfg.MySQLPassword = value
	case "MySQLDBNAME":
		cfg.MySQLDBNAME = value
	case "JWTKey":
		cfg.JWTKey = value
	case "JWTExpire":
		cfg.JWTExpire = value
	case "ENV":
		cfg.ENV = value
	}
}
