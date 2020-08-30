package config

type ConfigServerSt struct {
	Name    string
	Author  string
	Version string
	Port    uint32
}

func GetServerConfig(key string) ConfigServerSt {
	var conf ConfigServerSt
	conf.Name = GetString(key + "." + "name")
	conf.Author = GetString(key + "." + "author")
	conf.Version = GetString(key + "." + "version")
	conf.Port = GetUInt32(key + "." + "port")
	return conf
}

type ConfigLog struct {
	Level string
	Path  string
}

func GetLogConfig(key string) ConfigLog {
	var conf ConfigLog
	conf.Level = GetString(key + "." + "level")
	conf.Path = GetString(key + "." + "path")
	return conf
}

type ConfigRedis struct {
	Addr     string
	Password string
}

func GetRedisConfig(key string) ConfigRedis {
	var conf ConfigRedis
	conf.Addr = GetString(key + "." + "addr")
	conf.Password = GetString(key + "." + "password")
	return conf
}

type ConfigMysql struct {
	Addr     string
	User     string
	Password string
}

func GetMysqlConfig(key string) ConfigMysql {
	var conf ConfigMysql
	conf.Addr = GetString(key + "." + "addr")
	conf.User = GetString(key + "." + "user")
	conf.Password = GetString(key + "." + "password")
	return conf
}
