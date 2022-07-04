package another

import (
	"flag"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/url"
)

type MyURL url.URL

type Config struct {
	Port        string `default:"8080" envconfig:"Port"`
	DbUrl       MyURL  `default:"postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable" envconfig:"Db_url"`
	JaegerUrl   MyURL  `default:"http://jaeger:16686" envconfig:"jaeger_url"`
	SentryUrl   MyURL  `default:"http://sentry:9000" envconfig:"sentry_url"`
	KafkaBroker string `default:"kafka:9092" envconfig:"kafka_broker"`
	SomeAppId   string `default:"testid" envconfig:"some_app_id"`
	SomeAppKey  string `default:"testkey" envconfig:"some_app_key"`
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

func GetConfigurationFromFlags() Config {
	var configuration Config
	var dbUrl, jaegerUrl, sentryUrl string

	flag.StringVar(&configuration.Port, "port", "8080", "port to connect to DB")
	flag.StringVar(&dbUrl, "db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "DB url")
	configuration.DbUrl.Decode(dbUrl)
	flag.StringVar(&jaegerUrl, "jaeger_url", "http://jaeger:16686", "port to connect to DB")
	configuration.JaegerUrl.Decode(jaegerUrl)
	flag.StringVar(&sentryUrl, "sentry_url", "http://sentry:9000", "port to connect to DB")
	configuration.JaegerUrl.Decode(sentryUrl)
	flag.StringVar(&configuration.KafkaBroker, "kafka_broker", "kafka:9092", "port to connect to DB")
	flag.StringVar(&configuration.SomeAppId, "some_app_id", "testid", "port to connect to DB")
	flag.StringVar(&configuration.SomeAppKey, "some_app_key", "testkey", "app_key to connect to DB")

	flag.Parse()

	return configuration
}

func GetConfigurationFromEnv() Config {
	var configuration Config

	err := envconfig.Process("myapp", &configuration)

	if err != nil {
		log.Fatal(err.Error())
	}

	format := "Port: %d\nDbUrl: %s\nJaegerUrl: %f\nSentryUrl: %s\nKafkaBroker: %s\nSomeAppId: %s\nSomeAppKey: %s\n"
	_, err = fmt.Printf(format, configuration.Port, configuration.DbUrl, configuration.JaegerUrl, configuration.SentryUrl, configuration.KafkaBroker, configuration.SomeAppId, configuration.SomeAppKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	return configuration
}
