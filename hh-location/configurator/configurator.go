package configurator

import (
	"github.com/tkanos/gonfig"
	"log"
	"path/filepath"
	"sync"
)

type Configuration struct {
	DbString string
}

var confInst *Configuration
var once sync.Once

func GetConfiguration() *Configuration {
	once.Do(func() {
		confInst = load()
	})
	return confInst
}

func load() *Configuration {
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
