package utils

import "github.com/spf13/viper"

/* store all Configuration off the application. */
/* the value are read by the viper from config file or environment variables */

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadCongig read the configuration from file or environment variables.
func LoadConfig( path string ) ( config Config, err error ){
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
