package config

type ConfigServer struct {
	Name    string
	Author  string
	Version string
	Port    uint32
}

func GetServerConfig(key string) ConfigServer {
	var conf ConfigServer
	conf.Name = GetString(key + "." + "name")
	conf.Author = GetString(key + "." + "author")
	conf.Version = GetString(key + "." + "version")
	conf.Port = GetUInt32(key + "." + "port")
	return conf
}

type ConfigLog struct {
	Level      string
	Path       string
	MaxSize    uint32 // MB
	MaxBackups uint32 //file count
	MaxAge     uint32 //days
	Compress   bool
}

func GetLogConfig(key string) ConfigLog {
	var conf ConfigLog
	conf.Level = GetString(key + "." + "level")
	conf.Path = GetString(key + "." + "path")
	conf.MaxSize = GetUInt32(key + "." + "maxSize")
	conf.MaxBackups = GetUInt32(key + "." + "maxBackups")
	conf.MaxAge = GetUInt32(key + "." + "maxAge")
	conf.Compress = GetBool(key + "." + "compress")
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