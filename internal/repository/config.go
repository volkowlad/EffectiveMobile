package repository

const (
	songsTable = "songs"
	infoTable  = "info"
	textTable  = "text"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewConfigDB(host, port, username, password, dbName, sslMode string) *Config {
	return &Config{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		DBName:   dbName,
		SSLMode:  sslMode,
	}
}
