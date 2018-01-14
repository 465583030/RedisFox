package dataprovider

import (
	"fmonitor/conf"
)

type DataProvider interface {
	SaveMemoryInfo(server string, used int, peak int) int64
	SaveInfoCommand(server string, info map[string]string) int64
	SaveMonitorCommand(server, command, argument, keyname string) int64
	/*GetInfo(server string) (map[string]interface{}, error)
	GetMemoryInfo(server, fromDate, toDate string) ([]map[string]interface{}, error)*/
}

func NewProvider(config *conf.Config) DataProvider {
	if config.Datatype == "sqlite" {
		return NewSqliteProvide(config.Datapath)
	}
	return nil
}
