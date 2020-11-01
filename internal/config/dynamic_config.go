package config

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"time"
)

var GlobalDynamicConfig *DynamicConfig

type DynamicConfig struct {
	DBContextQueryTimeout time.Duration
	MaxRequestBodySize int64
}

// LoadFromFile load config from file
func (r *DynamicConfig) LoadFromFile(filename string) error {
	configFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	r.SetDefaultValue()
	err = xml.Unmarshal(configFile, r)
	if err != nil {
		return err
	}
	return nil
}

// Maintain load config periodically
func (r *DynamicConfig) Maintain(filename string, interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		for {
			<-ticker.C
			newConf := &DynamicConfig{}
			fmt.Printf("reload_dyanmic_config|trigger=ticker")
			if err := newConf.LoadFromFile(filename); err != nil {
				fmt.Errorf("reload_search_dynamic_failed|file=%s,error=%s", filename, err)
			} else {
				*r = *newConf
				r.SetDefaultValue()
				fmt.Printf("reload_search_dynamic_ok|sc=%v", *r)
			}
		}
	}()
}

func (globalcli *DynamicConfig) SetDefaultValue() {
	if globalcli.DBContextQueryTimeout <= time.Duration(0) {
		globalcli.DBContextQueryTimeout = time.Duration(time.Second / 2)
	}
	if globalcli.MaxRequestBodySize <= 0 {
		globalcli.MaxRequestBodySize = 5242880 //5MB
	}
}
