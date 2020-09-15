package config

import "github.com/spf13/viper"

//server
type ConfigServer struct {
	Name         string
	Author       string
	Version      string
	Port         uint32
	NetTimeOutMs uint32 //网络耗时，即client到server入口的网络耗时，超过这个时间请求直接返回，默认5000ms
}

func GetServerConfig(key string) ConfigServer {
	var conf ConfigServer
	conf.Name = GetString(key + "." + "name")
	conf.Author = GetString(key + "." + "author")
	conf.Version = GetString(key + "." + "version")
	conf.Port = GetUInt32(key + "." + "port")
	conf.NetTimeOutMs = GetUInt32(key + "." + "netTimeOutMs")
	return conf
}

//log
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

//redis
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

//mysql
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

//etcd
type ConfigEtcd struct {
	Addrs []string
}

func GetEtcdConfig(key string) ConfigEtcd {
	var conf ConfigEtcd
	conf.Addrs = GetStringSlice(key + "." + "addrs")
	return conf
}

//http client
type ConfigHttp struct {
	Addr           string //http addr
	TimeOutMs      uint32 //请求超时时间，单位毫秒
	Method         string //"POST" or "GET"
	ContentType    string //如"application/json; charset=utf-8"
	Scheme         string //"http" or "https"
	Host           string //http host,不填为无
	Proxy          string //代理，不填为无
	CertFilePath   string //https cert文件路径
	KeyFilePath    string //https key文件路径
	RootCaFilePath string //https ca filt
}

func GetHttpConfig(key string) ConfigHttp {
	viper.SetDefault(key+"."+"timeOutMs", 5000)
	viper.SetDefault(key+"."+"method", "GET")
	viper.SetDefault(key+"."+"scheme", "http")

	var conf ConfigHttp
	conf.Addr = GetString(key + "." + "addr")
	conf.TimeOutMs = GetUInt32(key + "." + "timeOutMs")
	conf.Method = GetString(key + "." + "method")
	conf.ContentType = GetString(key + "." + "contentType")
	conf.Scheme = GetString(key + "." + "scheme")
	conf.Host = GetString(key + "." + "host")
	conf.Proxy = GetString(key + "." + "proxy")
	conf.CertFilePath = GetString(key + "." + "certFilePath")
	conf.KeyFilePath = GetString(key + "." + "keyFilePath")
	conf.RootCaFilePath = GetString(key + "." + "rootCaFilePath")

	return conf
}

//grpc client
type ConfigGrpc struct {
	Name      string //http name
	TimeOutMs uint32 //请求超时时间，单位毫秒
}

func GetGrpcConfig(key string) ConfigGrpc {
	var conf ConfigGrpc
	conf.Name = GetString(key + "." + "name")
	conf.TimeOutMs = GetUInt32(key + "." + "timeOutMs")

	return conf
}

//net client
type ConfigNet struct {
	Addr      string //addr
	TimeOutMs uint32 //请求超时时间，单位毫秒
}

func GetNetConfig(key string) ConfigGrpc {
	var conf ConfigGrpc
	conf.Name = GetString(key + "." + "addr")
	conf.TimeOutMs = GetUInt32(key + "." + "timeOutMs")

	return conf
}
