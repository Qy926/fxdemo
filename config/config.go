// lib/config/config.go
package config

type httpConfig struct {
	ListenAddress string
}

type dbConfig struct {
	URL string
}

type Config struct {
	HTTP httpConfig
	DB   dbConfig
}
