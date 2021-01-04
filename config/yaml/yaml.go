package yaml

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

/**
 * 加载yaml配置文件
 * param: v, 接收配置参数的结构体指针，v需要自己初始化
 * param: configFile, 配置文件的存储路径
 * return：error, 如果成功返回nil, 否则返回错误信息
 */
func Load(v interface{}, configFile string) error {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("file=%s,err=%s", configFile, err.Error())
	}
	return yaml.Unmarshal(b, v)
}
