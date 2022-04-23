package config

type databaseConfig struct {
	DbHost     string
	DbPort     int
	DbUser     string
	DbName     string
	DbPassword string
	DbSchema   string
}

func newDatabaseConfig() databaseConfig {
	return databaseConfig{
		DbHost:     ReadEnvString("DB_HOST"),
		DbName:     ReadEnvString("DB_NAME"),
		DbPort:     ReadEnvInt("DB_PORT"),
		DbUser:     ReadEnvString("DB_USER"),
		DbPassword: ReadEnvString("DB_PASSWORD"),
		DbSchema:   ReadEnvString("DB_SCHEMA"),
	}
}
