package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
	"github.com/spf13/viper"
)

type DataConnection struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Schema   string
}

type AppConfigs struct {
	SqlServer DataConnection `yaml:"sqlServer"`
}



func GetConfigs(configuration *AppConfigs, configPath string) {
	if _, err := os.Stat(configPath); err == nil {
		content, err := os.ReadFile(configPath)
		if err != nil {
			log.Fatal("Error opening config file: ", err)
		}
		err = yaml.Unmarshal(content, &configuration)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}
	} else {
		log.Fatal("Config file not found: ", err)
	}
}

func LoadConfigs(configuration *AppConfigs, configPath string){

}


func main() {
	var configs AppConfigs
	var viperConifgs AppConfigs
	var configPath = "./configs/configs.yaml"
	GetConfigs(&configs, configPath)

	fmt.Printf("Configs \n Username: %s\n Password: %s\n Server: %s\n Host: %s\n Port: %d", configs.SqlServer.Username, configs.SqlServer.Password, configs.SqlServer.Schema, configs.SqlServer.Host, configs.SqlServer.Port)
}
