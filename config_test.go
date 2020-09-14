package config

import (
	"testing"
)

func TestConfigLoad(t *testing.T) {
	InitConfig()

	serverName := GetString("server.name")
	t.Logf("server.name: %s", serverName)
	serverName1 := GetString("server.name1")
	t.Logf("server.name: %s", serverName1)

	t.Logf("server: %+v", GetServerConfig("server"))
	t.Logf("log: %+v", GetLogConfig("log"))
	t.Logf("redis_test: %+v", GetRedisConfig("redis_test"))
	t.Logf("mysql_test: %+v", GetMysqlConfig("mysql_test"))
	t.Logf("etcd_test: %+v", GetEtcdConfig("etcd"))
	t.Logf("http_test: %+v", GetHttpConfig("http_test"))

}
