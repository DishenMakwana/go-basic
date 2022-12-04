package main

import (
	"encoding/json"
	"github.com/dishenmakwana/database/dbtools"
	"github.com/dishenmakwana/rest/restlayer"
	"log"
	"os"
)

type Configuration struct {
	DriverName     string `json:"driverName"`
	DataSourceName string `json:"dataSourceName"`
}

func main() {
	file, err := os.Open("database/configuration/config.json")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	conf := new(Configuration)

	json.NewDecoder(file).Decode(conf)

	dbtools.DBInitializer(conf.DriverName, conf.DataSourceName)

	restlayer.RestStart("127.0.0.1:8080")
}
