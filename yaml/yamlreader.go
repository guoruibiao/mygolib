package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//DB配置
type Db struct {
	Dbname  string `yaml:"dbname"`
	Charset string `yaml:"charset"`
	Index   []int  `yaml:"index"`
}

//Config
type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Db   Db     `yaml:"db"`
}

/* 有点奇怪，属性首字母小写就无法访问了
 * 翻了翻网上的解答：
 * - struct变量成员名小写时不符合golang的导出规则, 导致json.Marshal(其使用了reflect)无法反射到内容.
 * - struct内的成员变量小写就只有当前包内可以访问。如果把json.Marshal(hp) 这个方法的重新实现在main方法下面，struct中的成员变量就可以小写了
 */
// Struct fields are only unmarshalled if they are exported (have an
// upper case first letter), and are unmarshalled using the field name
// lowercased as the default key. Custom keys may be defined via the
// "yaml" name in the field tag: the content preceding the first comma
// is used as the key, and the following comma-separated options are
// used to tweak the marshalling process (see Marshal).
// Conflicting names result in a runtime error.
// see more: https://github.com/go-yaml/yaml/yaml.go

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
