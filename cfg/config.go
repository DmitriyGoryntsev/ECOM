package cfg

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SslMode  string
}

func (c *Config) ConnectionString() string {
	return "host=" + c.Host + " port=" + c.Port + " user=" + c.User + " password=" + c.Password + " dbname=" + c.DBName + " sslmode=" + c.SslMode
}

func NewConfig(host, port, user, password, dbName, sslMode string) *Config {
	return &Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
		SslMode:  sslMode,
	}
}
