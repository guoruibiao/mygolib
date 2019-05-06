package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//Nginx nginx  配置
type Db struct {
	Dbname  string `yaml:"dbname"`
	Charset string `yaml:"charset"`
	Index   []int  `yaml:"index"`
}

//Config   系统配置配置
type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Db   Db     `yaml:"db"`
}

func main() {

	var setting Config
	config, err := ioutil.ReadFile("./myconf.yaml")
	if err != nil {
		fmt.Print(err)
	}
	yaml.Unmarshal(config, &setting)

	fmt.Println("GET db.dbname: ", setting.Db.Dbname)
	fmt.Println("DBINDEX: ", setting.Db.Index[1])

}
