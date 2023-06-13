package config

type Config struct {
	Port int `env:"SERVER_PORT" envDefault:"13005"`

	PgPort   int `env:"PG_PORT" envDefault:"5458"`
	PgHost   int `env:"PG_HOST" envDefault:"0.0.0.0"`
	PgDBName int `env:"PG_DB_NAME" envDefault:"db"`
	PgUser   int `env:"PG_USER" envDefault:"db"`
	PgPwd    int `env:"PG_PWD" envDefault:"db"`
}
