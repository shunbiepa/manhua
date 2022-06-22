package config

import (
	"io/ioutil"
	"log"
	"sync"

	"gopkg.in/yaml.v2"
)

var once sync.Once
var configInfo = new(Specification)

type Specification struct {
	DB struct {
		Typ   string `yaml:"type"`
		DSN   string `yaml:"dsn"`
		Debug bool
	}
	Cache struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
	}
	MQType string `yaml:"MQType"`
	Kafka  struct {
		Addr string `yaml:"addr"`
	}
	Zap struct {
		Level        string `yaml:"level"`          // 级别
		Format       string `yaml:"format"`         // 输出
		Prefix       string `yaml:"prefix"`         // 日志前缀
		Director     string `yaml:"director"`       // 日志文件夹
		LogInConsole bool   `yaml:"log-in-console"` // 输出控制台
	}
}

func Get(fileName string) (*Specification, error) {
	once.Do(func() {
		if fileName == "" {
			fileName = "./conf/conf.yaml"
		}
		err := Refresh(fileName)
		if err != nil {
			log.Fatalln("read conf file: ", err)
		}
	})
	return configInfo, nil
}

func Refresh(fileName string) error {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(b, configInfo); err != nil {
		return err
	}
	return nil
}
