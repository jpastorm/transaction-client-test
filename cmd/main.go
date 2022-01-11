package main

import (
	"fmt"
	"log"

	"github.com/jpastorm/transaction-client-test/infrastructure/handler"
	"github.com/jpastorm/transaction-client-test/infrastructure/handler/response"
	"github.com/jpastorm/transaction-client-test/model"
)

func main() {
	config := newConfiguration("./configuration.json")
	api := newHTTP(config, response.HTTPErrorHandler)
	logger := newLogger(config.LogFolder, false)
	db, err := newDBConnection(config.Database)
	if err != nil {
		log.Fatal(err)
	}
	remoteConfig := newRemoteConfig(config.Database)

	handler.InitRoutes(model.RouterSpecification{
		Config:       config,
		Api:          api,
		Logger:       logger,
		DB:           db,
		RemoteConfig: remoteConfig,
	})

	port := fmt.Sprintf(":%d", config.ServerPort)
	if config.IsHttps {
		log.Fatal(api.StartTLS(port, config.CertPem, config.KeyPem))
	} else {
		log.Println("Starting service")
		log.Fatal(api.Start(port))
	}
}
