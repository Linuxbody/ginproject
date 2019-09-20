package utils

import (
	// Viper将读取key/value存储(如etcd或Consul)中的路径检索的配置字符串(如JSON,TOML,YAML或HCL)
	"github.com/spf13/viper"
	"time"
)

type Server struct {
	RunMode	string `json:"run_mode"`
	ServerAddr	string `json:"server_addrr"`
	ReadTimeout	time.Duration `json:"read_timeout"`
	WriteTimeout	time.Duration `json:"write_timeout"`
}

type App struct {
	TimeFormat	string `json:"time_format"`
	TokenTimeout	int64 `json:"token_timeout"`
	Schema	string	`json:"schema"`
	StaticBasePath	string	`json:"static_base_path"`
	UploadBasePath	string	`json:"upload_base_path"`
	ImageRelPath	string	`json:"image_rel_path"`
}

type DbMysql struct {
	Mode	string	`json:"mode"`
	Host	string	`json:"host"`
	Port	string	`json:"port"`
	DbUser	string	`json:"db_user"`
	DbPwd	string	`json:"db_pwd"`
	DbName	string	`json:"db_name"`
}

type Redis struct {
	RedisAddr	string	`json:"read_addr"`
	Pwd	string	`json:"pwd"`
	Db	int	`json:"db"`
	CacheTime	int	`json:"cache_time"`
}

var ServerInfo = &Server{}
var AppInfo = &App{}
var DbMysqlInfo = &DbMysql{}
var RedisInfo = &Redis{}

func init() {
	// viper.AddConfigPath("conf")
	viper.SetConfigFile("conf/config.yml")
	
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	ServerInfo.RunMode = viper.GetString("server.runMode")
	ServerInfo.ServerAddr = viper.GetString("server.serverAddr")
	ServerInfo.ReadTimeout = time.Duration(viper.GetInt("server.readTimeout")) * time.Second
	ServerInfo.WriteTimeout = time.Duration(viper.GetInt("server.writeTimeout")) * time.Second

	AppInfo.TimeFormat = viper.GetString("app.timeFormat")
	AppInfo.TokenTimeout = viper.GetInt64("app.tokenTimeout")
	AppInfo.Schema = viper.GetString("app.schema")
	AppInfo.StaticBasePath = viper.GetString("app.staticBasePath")
	AppInfo.UploadBasePath = viper.GetString("app.uploadBasePath")
	AppInfo.ImageRelPath = viper.GetString("app.imageRelPath")

	DbMysqlInfo.Mode = viper.GetString("dbmysql.mode")
	DbMysqlInfo.Host = viper.GetString("dbmysql.host")
	DbMysqlInfo.Port = viper.GetString("dbmysql.port")
	DbMysqlInfo.DbUser = viper.GetString("dbmysql.dbUser")
	DbMysqlInfo.DbPwd = viper.GetString("dbmysql.dbPwd")
	DbMysqlInfo.DbName = viper.GetString("dbmysql.dbName")

	RedisInfo.RedisAddr = viper.GetString("redis.redisAddr")
	RedisInfo.Pwd = viper.GetString("redis.pwd")
	RedisInfo.Db = viper.GetInt("redis.db")
	RedisInfo.CacheTime = viper.GetInt("redis.cacheTime")

}