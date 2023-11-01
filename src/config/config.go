package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = &Config{}

type Config struct {
	Server ServerConf
	Mysql  MysqlConf
	Log    LogConf
	JWT    JWTConf
	Redis  RedisConf
	Mail   MailConf
}

type ServerConf struct {
	Name          string `json:"name"`
	Listen        string `json:"listen"`
	Mode          string `json:"mode"`
	Env           string `json:"env"`
	Lang          string `json:"lang"`
	CoroutinePoll int    `json:"coroutinePoll"`
	Node          string `json:"node"`
	ServiceOpen   bool   `json:"serviceOpen"`
	GrpcListen    string `json:"grpcListen"`
	FilePath      string `json:"filePath"`
}

type JWTConf struct {
	Secret string `json:"secret"`
	Ttl    int64  `json:"ttl"`
}

type MysqlConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Charset  string `json:"charset"`
}

type LogConf struct {
	Level     string `json:"level"`
	Type      string `json:"type"`
	FileName  string `json:"filename"`
	MaxSize   int    `json:"maxSize"`
	MaxBackup int    `json:"maxBackup"`
	MaxAge    int    `json:"maxAge"`
	Compress  bool   `json:"compress"`
}

type RedisConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
	Poll     int    `json:"poll"`
	Conn     int    `json:"conn"`
}

type MailConf struct {
	Driver     string `json:"driver"`
	Host       string `json:"host"`
	Name       string `json:"name"`
	Port       int    `json:"port"`
	Password   string `json:"password"`
	Encryption string `json:"encryption"`
	FromName   string `json:"fromName"`
}

func InitConfig(path string) *Config {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		err = viper.Unmarshal(&Conf)
		if err != nil {
			panic(err)
		}
	})

	return Conf
}

func IsLocal() bool {
	return Conf.Log.Level == "local"
}

func IsProduction() bool {
	return Conf.Log.Level == "production"
}

func IsTesting() bool {
	return Conf.Log.Level == "testing"
}
