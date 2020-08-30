package config

import (
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	//log 默认值
	viper.SetDefault("log.level", "debug")
	viper.SetDefault("log.path", "../log/")
	viper.SetDefault("log.maxSize", 100)
	viper.SetDefault("log.maxBackups", 20)
	viper.SetDefault("log.maxAge", 30)
	viper.SetDefault("log.compress", false)

	viper.SetConfigName("config")   //设置文件名，后缀会自动识别，建议yml文件
	viper.AddConfigPath("../conf/") //相对可执行文件的路径
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Init Config File Err")
	}
}

func AllKeys() []string {
	return viper.AllKeys()
}

func Get(key string) interface{} {
	return viper.Get(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetInt32(key string) int32 {
	return viper.GetInt32(key)
}

func GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func GetUInt(key string) uint {
	return viper.GetUint(key)
}

func GetUInt32(key string) uint32 {
	return viper.GetUint32(key)
}

func GetUInt64(key string) uint64 {
	return viper.GetUint64(key)
}

func GetIntSlice(key string) []int {
	return viper.GetIntSlice(key)
}

func GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

func GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}

func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}
