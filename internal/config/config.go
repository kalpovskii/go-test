package config

type Config struct {
	Port int `env:"SERVER_PORT" envDefault:"13005"`

	PgPort   string `env:"PG_PORT" envDefault:"5432"`
	PgHost   string `env:"PG_HOST" envDefault:"192.168.0.10"`
	PgDBName string `env:"PG_DB_NAME" envDefault:"db"`
	PgUser   string `env:"PG_USER" envDefault:"db"`
	PgPwd    string `env:"PG_PWD" envDefault:"db"`
}
