package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Config is main cfg
type Config struct {
	Database []Db `json:"db"`
}

//Db is the configuration for connect to database
type Db struct {
	Type     string `json:"type"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Host     string `json:"host"`
	DBName   string `json:"database"`
}

var config Config

func getConfig() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &config)
}

func connect() *sql.DB {
	db, err := sql.Open(config.Database[0].Type, config.Database[0].Login+":"+config.Database[0].Password+
		"@tcp("+config.Database[0].Host+")/"+config.Database[0].DBName)
	if err != nil {
		panic(err)
	}
	return db
}
