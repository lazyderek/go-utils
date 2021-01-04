package yaml

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {

	type Server struct {
		Port string `yaml:"port"`
		Name string `yaml:"name"`
	}

	type Config struct {
		Server Server `yaml:"server"`
	}

	path := "./config.yml"
	var conf Config
	err := Load(&conf, path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("server\n name:"+conf.Server.Name, "\n port:"+conf.Server.Port)

}
