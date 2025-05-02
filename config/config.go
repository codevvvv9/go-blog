package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

// TomlConfig 映射配置文件 config.toml，虽然 key写的是大小，但是"github.com/BurntSushi/toml"可以处理成小写
type TomlConfig struct {
	Viewer Viewer
	System SystemConfig
}
type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

var Cfg *TomlConfig

// 程序启动时就会执行 init方法
func init() {
	Cfg = new(TomlConfig)
	var err error
	// 给配置赋值
	// 记录当前的目录
	Cfg.System.CurrentDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	Cfg.System.AppName = "wu-go-blog"
	Cfg.System.Version = 1.0
	_, err = toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		panic(err)
	}
}
