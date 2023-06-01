package config

import "os"

type (
	// Configuration struct
	Configuration struct {
		ConfigDB
	}

	// DB Config
	ConfigDB struct {
		ConnString string
	}
)

func Set() *Configuration {
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
		userDB = "pub-sub"
	}

	pswDB := os.Getenv("PSW_DB")
	if pswDB == "" {
		pswDB = "qwerty"
	}

	nameDB := os.Getenv("NAME_DB")
	if nameDB == "" {
		nameDB = "pub-sub"
	}

	ssldb := os.Getenv("SSLDB")
	if ssldb == "" {
		ssldb = "disable"
	}

	connStr := "host=" + hostDB + " port=" + portDB + " user=" + userDB + " password=" + pswDB + " dbname=" + nameDB + " sslmode=" + ssldb

	return &Configuration{ConfigDB{ConnString: connStr}}
}
