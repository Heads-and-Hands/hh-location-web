package configurator

import (
	"github.com/tkanos/gonfig"
	"log"
	"path/filepath"
)

type Configuration struct {
	DbString 			string
}

var confInst Configuration = Configuration{}

func GetConfiguration() Configuration {
	if (Configuration{}) == confInst {
		confInst = load()
	}
	//return confInst
	return load()
}

func load() Configuration {
	result := Configuration{}
	absPath, _ := filepath.Abs("./bin/config/config.json")
	err := gonfig.GetConf(absPath, &result)
	if err != nil {
		log.Println(err)
		return Configuration{}
	} else {
		return result
	}
}

