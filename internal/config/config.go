package config

import (
	"fmt"
	"os"
)

type (
	// Configuration struct
	Configuration struct {
		ConfigDB ConfigDB
	}

	// DB Config
	ConfigDB struct {
		ConnString string
	}
)

func Set() *Configuration {
	cfgDB := setDB()

	return &Configuration{
		ConfigDB: cfgDB,
	}
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
