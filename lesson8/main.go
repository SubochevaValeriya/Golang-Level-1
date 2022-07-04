package main

import (
	"fmt"
	another "pwd/anotherPackage"
)

func main() {

	// go run main.go --Port= --Db_url= -Jaeger_url= --MyURL= --Sentry_url= --Kafka_broker= --Some_app_id= --Some_app_key=

	fmt.Println(another.GetConfiguration())

}
