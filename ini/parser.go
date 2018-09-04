package ini

import (
	"io/ioutil"
	"fmt"
	"strings"
)

// 读取ini文件到对应的字典中
func read_ini_file(filepath string) (map[string]map[string]string, error) {
	ret := make(map[string]map[string]string)
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	content := string(bytes)
	blocks := strings.Split(content, "[")
	// 解析block
	for _, block := range blocks {
		// 第一行肯定是member的名字
		lines := strings.Split(block, "\n")
		member := strings.TrimRight(lines[0], "]")
		tmpmember := make(map[string]string)
		// 解析每一个配置项
		for _, line := range lines[1:] {
			line = strings.Trim(strings.Trim(line, " "), "\n")
			if line == "" {
				continue
			}
			if strings.IndexRune(line, ';') == 0 {
				continue
			}
			item := strings.Split(line, "=")
			key, value := item[0], item[1]
			tmpmember[key] = value
		}
		ret[member] = tmpmember
	}
	return ret, nil
	//return make(map[string]map[string]string)
}


func get_ini_member(filepath string, member string) (map[string]string, error) {
	ret, err := read_ini_file(filepath)
	if err != nil {
		return _, err
	}
	return ret[member], nil
	
}

func main() {
	filepath := "./config.ini"
	//ret := read_ini_file(filepath)
	//fmt.Println(ret)
	user, _ := get_ini_member(filepath, "User")
	fmt.Println(user["name"])
	deploy, _ := get_ini_member(filepath, "Deploy")
	fmt.Println(deploy["debug"])
}
