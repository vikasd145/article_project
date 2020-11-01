package config

import "time"

var GlobalDynamicConfig *DynamicConfig

type DynamicConfig struct {
	DBContextQueryTimeout time.Duration
}

func (globalcli *DynamicConfig) SetDefaultValue() {
	if globalcli.DBContextQueryTimeout <= time.Duration(0) {
		globalcli.DBContextQueryTimeout = time.Duration(time.Second / 2)
	}
}
