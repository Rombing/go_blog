package core

import (
	"gopkg.in/yaml.v3"
	"log"
	"server/config"
	"server/utils/my_yaml"
)

func InitConfig() *config.Config {
	Config := &config.Config{}
	yamlConf, err := my_yaml.LoadYAML()
	if err != nil {
		log.Fatalf("my_yaml load config error:%v", err)
	}
	err = yaml.Unmarshal(yamlConf, Config)
	if err != nil {
		log.Fatalf("yaml unmarshal config error:%v", err)
	}
	return Config
}
