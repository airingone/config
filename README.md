#配置文件组建
## 1.组件描述
config用于加载与读取yml配置文件，默认封装ConfigServer，ConfigLog，ConfigRedis等等配置结构。
配置文件格式请参考[config.yml](https://github.com/airingone/config/blob/master/config.yml)。

## 2.如何使用
```
import (
    "github.com/airingone/config"
)

func main() {
    config.InitConfig() //进程启动时调用一次初始化配置文件，配置文件名为config.yml，目录路径为../conf/或./

    serverName := config.GetString("server.name")      //获取指定某个配置项的值
    configRedis := config.GetRedisConfig("redis_test") //获取redis配置结构
}
```
更多使用请参考[测试例子](https://github.com/airingone/config/blob/master/config_test.go)
