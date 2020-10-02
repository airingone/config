package config

import "github.com/spf13/viper"

//定义基础的config，为了更方便在各组件模块使用，这里可以持续增加

//server config
type ConfigServer struct {
	Name          string //服务名，各模块需要取不一样的名字，如果使用etcd发现服务则使用服务名，建议XXX_name
	Author        string //服务负责人
	Version       string //版本
	Port          uint32 //服务端口，可http/tcp/udp
	NetTimeOutMs  uint32 //网络超时,单位毫秒，默认5000
	IdleTimeoutMs uint32 //tcp服务时的空闲时间上限，单位秒，默认300
	CapacityPool  uint32 //协程池最大协程数，默认10000
	CapacityLimit uint32 //限频值，每秒最大请求量为该值，默认10000
}

//get server config
func GetServerConfig(key string) ConfigServer {
	viper.SetDefault("server.netTimeOutMs", 5000)
	viper.SetDefault("server.idleTimeoutMs", 300000)
	viper.SetDefault(key+"."+"capacityPool", 10000)
	viper.SetDefault(key+"."+"capacityLimit", 10000)

	var conf ConfigServer
	conf.Name = GetString(key + "." + "name")
	conf.Author = GetString(key + "." + "author")
	conf.Version = GetString(key + "." + "version")
	conf.Port = GetUInt32(key + "." + "port")
	conf.NetTimeOutMs = GetUInt32(key + "." + "netTimeOutMs")
	conf.IdleTimeoutMs = GetUInt32(key + "." + "idleTimeoutMs")
	conf.CapacityPool = GetUInt32(key + "." + "capacityPool")
	conf.CapacityLimit = GetUInt32(key + "." + "capacityLimit")
	return conf
}

//log config
type ConfigLog struct {
	Level      string //日志等级，支持error,debug,info,warn,trace，默认debug
	Path       string //日志路径，默认../log/
	MaxSize    uint32 //日志文件最大大小，单位MB，默认100
	MaxBackups uint32 //文件数，默认20
	MaxAge     uint32 //日志过期时间，单位天，默认30
	Compress   bool   //是否压缩，默认false
}

//get log config
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

//redis config
type ConfigRedis struct {
	Addr     string //地址，ip:port
	Password string //密码
}

//get redis config
func GetRedisConfig(key string) ConfigRedis {
	var conf ConfigRedis
	conf.Addr = GetString(key + "." + "addr")
	conf.Password = GetString(key + "." + "password")
	return conf
}

//mysql config
type ConfigMysql struct {
	Addr         string //地址，如"root:test.2020@tcp(127.0.0.1:3306)/air_test?charset=utf8"
	MaxIdleConns uint32 //最大空闲连接
	MaxOpenConns uint32 //最大打开连接数
}

//get mysql config
func GetMysqlConfig(key string) ConfigMysql {
	viper.SetDefault(key+"."+"maxIdleConns", 10)
	viper.SetDefault(key+"."+"maxOpenConns", 300)
	var conf ConfigMysql
	conf.Addr = GetString(key + "." + "addr")
	conf.MaxIdleConns = GetUInt32(key + "." + "maxIdleConns")
	conf.MaxOpenConns = GetUInt32(key + "." + "maxOpenConns")
	return conf
}

//etcd config
type ConfigEtcd struct {
	Addrs []string //etcd集群地址
}

//get etcd config
func GetEtcdConfig(key string) ConfigEtcd {
	var conf ConfigEtcd
	conf.Addrs = GetStringSlice(key + "." + "addrs")
	return conf
}

//http client config
type ConfigHttp struct {
	Addr           string //地址，可为: ip:ip:port, etcd:servername, url:www.xxx.com
	TimeOutMs      uint32 //请求超时时间，单位毫秒
	Method         string //http method，可为POST,GET
	ContentType    string //请求数据格式，如"application/json; charset=utf-8"
	Scheme         string //http scheme,可为http,https
	Host           string //host,如果填了则会使用host，默认为空不使用
	Proxy          string //proxy,如果填了则会使用proxy，默认为空不使用
	CertFilePath   string //https证书cert，如果是scheme=https则需要填
	KeyFilePath    string //https证书key，如果是scheme=https则需要填
	RootCaFilePath string //https证书rootCa，如果是scheme=https则需要填
}

//get http client config
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

//grpc client config
type ConfigGrpc struct {
	Name      string //请求的服务名，根据服务名注册etcd
	TimeOutMs uint32 //请求超时时间，单位毫秒
}

//get grpc client config
func GetGrpcConfig(key string) ConfigGrpc {
	var conf ConfigGrpc
	conf.Name = GetString(key + "." + "name")
	conf.TimeOutMs = GetUInt32(key + "." + "timeOutMs")

	return conf
}

//net client config
type ConfigNet struct {
	ServerName string //请求的服务名
	Addr       string //地址，可为: ip:ip:port, etcd:servername
	TimeOutMs  uint32 //请求耗时
}

//get net client config
func GetNetConfig(key string) ConfigNet {
	var conf ConfigNet
	conf.ServerName = GetString(key + "." + "serverName")
	conf.Addr = GetString(key + "." + "addr")
	conf.TimeOutMs = GetUInt32(key + "." + "timeOutMs")

	return conf
}

//es client config
type ConfigEs struct {
	Addr      string //地址，可为: "http://127.0.0.1:9200"
	UserName  string //用户名，如果无用户名与密码则直接为空即可
	Password  string //用户密码，如果无用户名与密码则直接为空即可
	TimeOutMs uint32 //请求耗时
}

//get es client config
func GetEsConfig(key string) ConfigEs {
	var conf ConfigEs
	conf.Addr = GetString(key + "." + "addr")
	conf.UserName = GetString(key + "." + "userName")
	conf.Password = GetString(key + "." + "password")
	conf.TimeOutMs = GetUInt32(key + "." + "timeOutMs")

	return conf
}
