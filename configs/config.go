package configs

import (
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	DBDriver           string `mapstructure:"DB_DRIVER"`
	DBHost             string `mapstructure:"DB_HOST"`
	DBPorta            string `mapstructure:"DB_PORTA"`
	DBUsuario          string `mapstructure:"DB_USUARIO"`
	DBSenha            string `mapstructure:"DB_SENHA"`
	DBNome             string `mapstructure:"DB_NOME"`
	WebServerPorta     string `mapstructure:"WEB_SERVER_PORTA"`
	GRPCServerPorta    string `mapstructure:"GRPC_SERVER_PORTA"`
	GraphQLServerPorta string `mapstructure:"GRAPHQL_SERVER_PORTA"`
}

func LoadConfig(path string) (*conf, error) {
	// cfg = &conf{}

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(path + "/.env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		panic(err)
	}

	return cfg, nil
}
