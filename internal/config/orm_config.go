package config

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

//OrmConfig Keeping all the orm config
type OrmConfig struct {
	MasterDSN    string `xml:"Master" json:"master_dsn,omitempty"`
	DBName       string `xml:"Name" json:"db_name,omitempty"`
	MaxConn      int    `xml:"MaxConn" json:"max_conn,omitempty"`
	MaxIdle      int    `xml:"MaxIdle" json:"max_idle,omitempty"`
	MaxConnSlave int    `xml:"MaxConnSlave" json:"max_conn_slave,omitempty"`
	MaxIdleSlave int    `xml:"MaxIdleSlave" json:"max_idle_slave,omitempty"`
}

type OrmContainer struct {
	Ormconfigs OrmConfig `xml:"Instance"`
}

// LoadFromFile ...
func (c *OrmContainer) LoadFromFile(filename string) (err error) {
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
