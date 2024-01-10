package configuration

import "strings"

const (
	DBTYPE_MYSQL  = "mysql"
	DBTYPE_SQLITE = "sqlite"
	DBTYPE_MEMORY = "memory"
)

type Configuration struct {
	Webserver struct {
		ListenAt string
		UseHttps bool
		Cert     string
	}
	Database struct {
		ConnectString string
		DBType        string
	}
}

func DetermineDBType(source string) string {
	switch strings.ToLower(source) {
	case "mysql":
		return DBTYPE_MYSQL
	case "sqlite":
		return DBTYPE_SQLITE
	default:
		return DBTYPE_MEMORY
	}
}
