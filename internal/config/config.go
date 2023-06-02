package config

import (
	"fmt"
	"os"
)

const SECRETKEY = "password"

type (
	// Configuration struct
	Configuration struct {
		HTTPServer HTTPServerConfiguration
		ConfigDB   ConfigDB
	}

	// HTTP server port
	HTTPServerConfiguration struct {
		Port string
	}

	// DB Config
	ConfigDB struct {
		ConnString string
	}
)

func Set() *Configuration {
	httpServerConf := setHTTP()
	cfgDB := setDB()

	return &Configuration{
		HTTPServer: httpServerConf,
		ConfigDB:   cfgDB,
	}
}

func setHTTP() HTTPServerConfiguration {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = ":8080"
	}

	return HTTPServerConfiguration{Port: httpPort}
}

func setDB() ConfigDB {
	hostDB := os.Getenv("HOST_DB")
	if hostDB == "" {
		hostDB = "localhost"
	}

	portDB := os.Getenv("PORT_DB")
	if portDB == "" {
		portDB = "5432"
	}

	userDB := os.Getenv("USER_DB")
	if userDB == "" {
		userDB = "admin"
	}

	pswDB := os.Getenv("PSW_DB")
	if pswDB == "" {
		pswDB = "password"
	}

	nameDB := os.Getenv("NAME_DB")
	if nameDB == "" {
		nameDB = "chartDB"
	}

	ssldb := os.Getenv("SSLDB")
	if ssldb == "" {
		ssldb = "disable"
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", hostDB, portDB, userDB, pswDB, nameDB, ssldb)

	return ConfigDB{ConnString: connStr}
}
