package configurator

import (
	"github.com/tkanos/gonfig"
	"log"
	"path/filepath"
	"sync"
	"os"
)

type Configuration struct {
	DbString string
}

var confInst *Configuration
var once sync.Once

func GetConfiguration() *Configuration {
	once.Do(func() {
		confInst = loadFromEnv()
	})
	return confInst
}

func loadFromEnv() *Configuration {
	// "root:rootroot@tcp(localhost:3306)/hh-location?parseTime=true"
	result := &Configuration{}

	result.DbString = os.Getenv("MYSQL_USER") + ":" +
		os.Getenv("MYSQL_PASSWORD") + "@tcp(" +
		os.Getenv("MYSQL_HOST") + ")/" +
		os.Getenv("MYSQL_DATABASE") + "?parseTime=true"

	log.Println("DB: ")
	log.Println(result.DbString)

	return result
}

func loadFromJSON() *Configuration {
	result := &Configuration{}
	absPath, _ := filepath.Abs("./bin/config/config.json")
	err := gonfig.GetConf(absPath, result)
	if err != nil {
		log.Println(err)
		return &Configuration{}
	} else {
		return result
	}
}
