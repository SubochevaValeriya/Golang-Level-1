package readConfiguration

import (
	"encoding/json"
	"flag"
	"github.com/go-yaml/yaml"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/url"
	"strings"
)

// base structure of configuration with tags for env, json & yaml

type Config struct {
	Port        int    `default:"8080" envconfig:"Port" json:"port" yaml:"port"`
	DbUrl       MyURL  `default:"postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable" envconfig:"Db_url" json:"db_url" yaml:"db_url"`
	JaegerUrl   MyURL  `default:"http://jaeger:16686" envconfig:"jaeger_url" json:"jaeger_url" yaml:"jaeger_url"`
	SentryUrl   MyURL  `default:"http://sentry:9000" envconfig:"sentry_url" json:"sentry_url" yaml:"sentry_url"`
	KafkaBroker string `default:"kafka:9092" envconfig:"kafka_broker" json:"kafka_broker" yaml:"kafka_broker"`
	SomeAppId   string `default:"testid" envconfig:"some_app_id" json:"some_app_id" yaml:"some_app_id"`
	SomeAppKey  string `default:"testkey" envconfig:"some_app_key" json:"some_app_key" yaml:"some_app_key"`
}

// our type for URL

type MyURL url.URL

// methods for type MyURL to decode and unmarchal with validation

func (u *MyURL) Decode(value string) error {
	_, err := url.Parse(value)
	if err == nil {
		r, _ := url.Parse(value)
		*u = *(*MyURL)(r)
	}

	return err
}

func (u *MyURL) UnmarshalJSON(value []byte) error {
	URLStrings := strings.Replace(string(value), "\"", "", -1)
	_, err := url.Parse(URLStrings)
	if err == nil {
		r, _ := url.Parse(URLStrings)
		*u = *(*MyURL)(r)
	}

	return err
}

func (u *MyURL) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var URLStrings string
	err := unmarshal(&URLStrings)
	_, err = url.Parse(URLStrings)
	if err == nil {
		r, _ := url.Parse(URLStrings)
		*u = *(*MyURL)(r)
	}

	return err
}

type Decoder interface {
	Decode(value string) error
}

type Unmarshaler interface {
	UnmarshalJSON([]byte) error
	UnmarshalYAML(unmarshal func(interface{}) error) error
}

// reading configuration from flags

func GetConfigurationFromFlags() Config {
	var configuration Config
	var dbUrl, jaegerUrl, sentryUrl string

	flag.IntVar(&configuration.Port, "port", 8080, "port to connect to DB")
	flag.StringVar(&dbUrl, "db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "DB url")
	configuration.DbUrl.Decode(dbUrl)
	flag.StringVar(&jaegerUrl, "jaeger_url", "http://jaeger:16686", "port to connect to DB")
	configuration.JaegerUrl.Decode(jaegerUrl)
	flag.StringVar(&sentryUrl, "sentry_url", "http://sentry:9000", "port to connect to DB")
	configuration.SentryUrl.Decode(sentryUrl)
	flag.StringVar(&configuration.KafkaBroker, "kafka_broker", "kafka:9092", "port to connect to DB")
	flag.StringVar(&configuration.SomeAppId, "some_app_id", "testid", "port to connect to DB")
	flag.StringVar(&configuration.SomeAppKey, "some_app_key", "testkey", "app_key to connect to DB")

	flag.Parse()

	return configuration
}

// reading configuration from Environmental Variables

func GetConfigurationFromEnv() Config {
	var configuration Config

	err := envconfig.Process("myapp", &configuration)

	if err != nil {
		log.Fatal(err.Error())
	}

	//format := "Port: %d\nDbUrl: %s\nJaegerUrl: %f\nSentryUrl: %s\nKafkaBroker: %s\nSomeAppId: %s\nSomeAppKey: %s\n"
	//_, err = fmt.Printf(format, configuration.Port, configuration.DbUrl, configuration.JaegerUrl, configuration.SentryUrl, configuration.KafkaBroker, configuration.SomeAppId, configuration.SomeAppKey)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}

	return configuration
}

// reading configuration from JSON file

func GetConfigurationFromJson(data []byte) Config {

	var configuration Config

	if err := json.Unmarshal(data, &configuration); err != nil {
		panic(err)
	}
	return configuration
}

// reading configuration from YAML file

func GetConfigurationFromYaml(data []byte) Config {
	var configuration Config

	if err := yaml.Unmarshal(data, &configuration); err != nil {
		panic(err)
	}
	return configuration
}
