package config

type Config struct {
	MongoConfig *MongoConfig
}

type MongoConfig struct {
	Username string
	Password string
	Database string
}

func GetConfig() *Config {
	return &Config{
		MongoConfig: &MongoConfig{
			Username: "Admin",
			Password: "Password1!",
			Database: "Authorization",
		},
	}
}
