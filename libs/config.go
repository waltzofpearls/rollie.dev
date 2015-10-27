package libs

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Env    string `json:"env"`
	Listen struct {
		Address string `json:"address"`
	} `json:"listen"`
	Github struct {
		Token    string `json:"token"`
		Username string `json:"username"`
	} `json:"github"`
}

func NewConfig() *Config {
	return &Config{}
}

func NewConfigFile(path string) *Config {
	c := NewConfig()
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)

	err := decoder.Decode(c)
	if err != nil {
		panic(fmt.Sprintf("Error parsing JSON config file: %s", err))
	}

	return c
}
