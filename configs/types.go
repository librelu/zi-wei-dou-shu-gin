package configs

import "github.com/gin-gonic/gin"

const (
	configNamePattern     = "%s-config"
	DefaultConfigYamlPath = "./configs/yaml"
	DefaultConfigType     = "yaml"
)

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Host     string
	Port     int32
	User     string
	Password string
	Database string
	SSLMode  string
}

// envMAp the env map reflect the gin mode
var envMap = map[string]bool{
	gin.DebugMode:   true,
	gin.ReleaseMode: true,
	gin.TestMode:    true,
}
