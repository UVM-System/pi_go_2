package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type capConfig struct {
	VideoId int `yaml:"videoid"`
	Prefix string `yaml:"prefix"`
}

var (
	Config conf
)

type conf struct {
	CapConfigs []capConfig `yaml:"caps"`
	DetectUrl string `yaml:"detecturl"`
	CapStartPin int `yaml:"capstartpin"`
	CapEndPin int `yaml:"capendpin"`
	Delay int `yaml:"delay"`
	TokenUrl string `yaml:"tokenurl"`
	MachineId string `yaml:"machineid"`
	Password string `yaml:"password"`
}

func init()  {
	Config.getConf()
}

func (c *conf) getConf() {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
