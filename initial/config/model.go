package config


// Config 全局配置示例
var Config config

//
const (
	serverName    = "AppContainerUploaderService"
	configVersion = "v1"
)

// config 全部配置
type config struct {
	// 应用配置功能
	App app `mapstructure:"app"`

	//数据库相关链接
	DB db `mapstructure:"db"`

	//存储目录
	Storage storage `mapstructure:"storage"`

}
// 应用Url配置
type app struct {
	Url string    `mapstructure:"url"`
	Address string   `mapstructure:"address"`
}

type db struct {
	// 连接字符串
	//Dsn string `mapstructure:"dsn"`
	//
	//// 密码
	//Pwd string `mapstructure:"pwd"`

	// 最大空闲连接数
	MaxIdleConn int `mapstructure:"maxidle_conn"`

	// 最大连接数
	MaxOpenConn int `mapstructure:"maxopen_conn"`
}

type storage struct {
	// 存储目录
	Dir    string  `mapstructure:"dir"`
}

