package config

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

var (
	Globalconf *Config
)

//Config Config of admin
type Config struct {
	//Ormconfigs db config
	Ormconfigs OrmConfig `xml:"Instance"`
	//RedisHost address of redis host running with port
	RedisHost string `xml:"redis"`
}

// LoadFromFile ...
func (c *Config) LoadFromFile(filename string) (err error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = xml.Unmarshal(file, c)
	if err != nil {
		fmt.Errorf("Error in unmarshalling error:%v", err)
		return err
	}
	return
}
