package config

import "github.com/spf13/viper"

type (
	Config struct {
		App      App      `mapstructure:"app"`
		Database Database `mapstructure:"database"`
	}

	App struct {
		LogLevel string `mapstructure:"log_level"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
	}

	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
		SSLMode  string `mapstructure:"sslmode"`
	}
)

func NewConfig(path string) *viper.Viper {
	vp := viper.New()
	vp.SetConfigFile(path)
	vp.AutomaticEnv()
	return vp
}

func Load(vp *viper.Viper) (*Config, error) {
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	var cfg Config
	if err := vp.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
