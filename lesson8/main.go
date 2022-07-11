package main

import (
	"fmt"
	"os"
	readConfiguration "pwd/anotherPackage"
)

func main() {

	//READ CONFIGURATION

	// USING FLAGS:
	//go run main.go --Port= --Db_url= -Jaeger_url= --MyURL= --Sentry_url= --Kafka_broker= --Some_app_id= --Some_app_key=
	fmt.Printf("Configuration from flags: %v\n", readConfiguration.GetConfigurationFromFlags())

	// USING ENVIRONMENTAL VARIABLES
	// SET (Windows)/EXPORT (Linux) ENVs
	fmt.Printf("Configuration from env: %v\n", readConfiguration.GetConfigurationFromEnv())

	// FROM JSON:
	dataJson, err := os.ReadFile("configuration.json")

	if err == nil {
		fmt.Printf("Configuration from JSON: %v\n", readConfiguration.GetConfigurationFromJson(dataJson))
	} else {
		fmt.Println(err)
	}

	//FROM YAML:
	dataYaml, err := os.ReadFile("configuration.yaml")

	if err == nil {
		fmt.Printf("Configuration from YAML: %v\n", readConfiguration.GetConfigurationFromYaml(dataYaml))
	} else {
		fmt.Println(err)
	}
}
