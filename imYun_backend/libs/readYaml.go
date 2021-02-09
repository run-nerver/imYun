package libs

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type YConfig struct {
	Host string `json:"host"`
	Port string `json:"port" comment:"#端口"`
	DB   struct {
		Host     string `json:"host" comment:"Mysql地址"`
		Port     string `json:"port" comment:"Mysql端口"`
		UserName string `json:"userName" comment:"Mysql账号"`
		PassWord string `json:"passWord" comment:"Mysql密码"`
		Name     string `json:"name" comment:"数据库名称"`
	}
	Redis struct {
		Host     string `json:"host" comment:"Redis地址"`
		Port     string `json:"port" comment:"Redis端口"`
		PassWord string `json:"passWord" comment:"Redis密码"`
		DB0      int    `json:"db0"`
		DB1      int    `json:"db1"`
		DB2      int    `json:"db2"`
	}
	Wechat struct {
		AppID  string `json:"appId"`
		Secret string `json:"secret"`
	}
	Oss struct {
		Enable string `json:"enable"`

		AccessKeyID     string `json:"accessKeyId"`
		AccessKeySecret string `json:"accessKeySecret"`
		Endpoint        string `json:"endpoint"`
	}
}

var YamlConfig *YConfig

func (c *YConfig) GetConfig() *YConfig {
	yamlConfig, err := ioutil.ReadFile("../build/config.repo.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlConfig, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
