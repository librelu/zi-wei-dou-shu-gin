package configs

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/zi-wei-dou-shu-gin/utils/utilerrors"
)

func NewConfigs(configPath, configFileType string) (*Config, error) {
	initConfig(configFileType)
	ginMode, err := GetGINMode()
	if err != nil {
		return nil, utilerrors.Wrap(err, "get gin-mode error when new configs")
	}

	configFilename := GetConfigFilename(ginMode)
	if err := readConfig(configFilename, configPath); err != nil {
		return nil, utilerrors.Wrap(err, "reading config error when new configs")
	}

	dbConfig, err := NewDBConfig()
	if err != nil {
		return nil, utilerrors.Wrap(err, "getting DB Config failed when new configs")
	}

	return &Config{
		DB: dbConfig,
	}, nil
}

func ResetConfig() {
	viper.Reset()
}

func initConfig(configFileType string) {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("GIN")
	viper.BindEnv("MODE")
	viper.SetConfigType(configFileType)
}

func readConfig(configFilename, configPath string) error {
	viper.SetConfigName(configFilename)
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return utilerrors.Wrap(err, "read configs failed")
	}
	return nil
}

func GetGINMode() (string, error) {
	env := viper.GetString("mode")
	if _, ok := envMap[env]; !ok {
		return "", utilerrors.New(
			fmt.Sprintf("unknown env:%s", env),
		)
	}
	return env, nil
}

func GetConfigFilename(env string) string {
	return fmt.Sprintf(configNamePattern, env)
}
