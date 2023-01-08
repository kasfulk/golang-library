package config

type DatabaseEnvironment struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type RedisEnvirontment struct {
	Url      string
	Password string
}
