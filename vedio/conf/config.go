package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var ConfigYaml Config

type Config struct {
	Name string `yaml:"Name"`
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`

	Mysql struct {
		DataSource string `yaml:"DataSource"`
	} `yaml:"Mysql"`

	MyRedis struct {
		Host string `yaml:"Host"`
		Pass string `yaml:"Pass"`
		Type string `yaml:"Type"`
	} `yaml:"MyRedis"`

	Minio struct {
		Endpoint        string `yaml:"Endpoint"`
		AccessKeyID     string `yaml:"AccessKeyID"`
		SecretAccessKey string `yaml:"SecretAccessKey"`
	} `yaml:"Minio"`
}

func LoadConfig() error {
	File, err := ioutil.ReadFile("./conf/vedio.yaml")
	//File, err := ioutil.ReadFile("../conf/vedio.yaml")//配合miniInit_test.go测试文件使用
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(File, &ConfigYaml)
	if err != nil {
		return err
	}
	return nil
}
