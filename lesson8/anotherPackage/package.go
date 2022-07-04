package another

import (
	"flag"
	"net/url"
)

type MyURL url.URL

type Config struct {
	Port         string
	Db_url       MyURL
	Jaeger_url   MyURL
	Sentry_url   MyURL
	Kafka_broker string
	Some_app_id  string
	Some_app_key string
}

func (u *MyURL) Decode(value string) error {
	_, err := url.Parse(value)
	if err == nil {
		r, _ := url.Parse(value)
		*u = *(*MyURL)(r)
	}

	return err
}

type Decoder interface {
	Decode(value string) error
}

func GetConfiguration() Config {
	var configuration Config
	var dbUrl, jaegerUrl, sentryUrl string

	flag.StringVar(&configuration.Port, "port", "8080", "port to connect to DB")
	flag.StringVar(&dbUrl, "db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "DB url")
	configuration.Db_url.Decode(dbUrl)
	flag.StringVar(&jaegerUrl, "jaeger_url", "http://jaeger:16686", "port to connect to DB")
	configuration.Jaeger_url.Decode(jaegerUrl)
	flag.StringVar(&sentryUrl, "sentry_url", "http://sentry:9000", "port to connect to DB")
	configuration.Jaeger_url.Decode(sentryUrl)
	flag.StringVar(&configuration.Kafka_broker, "kafka_broker", "kafka:9092", "port to connect to DB")
	flag.StringVar(&configuration.Some_app_id, "some_app_id", "testid", "port to connect to DB")
	flag.StringVar(&configuration.Some_app_key, "some_app_key", "testkey", "app_key to connect to DB")

	flag.Parse()

	return configuration
}
