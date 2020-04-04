package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type TonWebServerConfig struct {
	PgHost        string `env:"PG_HOST" envDefault:"localhost"`
	PgPort        int32  `env:"PG_PORT" envDefault:"5432"`
	PgName        string `env:"PG_NAME,required"`
	PgUser        string `env:"PG_USER,required"`
	PgPwd         string `env:"PG_PWD,required"`
	RPCListenPort int32  `env:"RPC_LISTEN_PORT" envDefault:"5400"`
	WebListenPort int32  `env:"WEB_LISTEN_PORT" envDefault:"9999"`
	WebDomain     string `env:"WEB_DOMAIN" envDefault:"tonbet.io"`
	TonAPIHost    int32  `env:"TON_API_HOST,required"`
	TonAPIPort    int32  `env:"TON_API_PORT,required"`
}

func GetConfig() TonWebServerConfig {
	cfg := &TonWebServerConfig{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal("Cannot parse initial ENV vars: ", err)
	}
	return *cfg
}
