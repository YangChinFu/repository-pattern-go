package config

import (
	"github.com/spf13/viper"
)

//Config 設定檔struct
type Config struct {
	v *viper.Viper
}

// LoadConfigFromJSON 讀json設定檔
func LoadConfigFromJSON(c *Config) error {
	c.v = viper.New()
	// 設定檔名稱
	c.v.SetConfigName("config")
	// 設定的路徑
	c.v.AddConfigPath("./")
	// 設定檔的格式
	c.v.SetConfigType("json")
	if err := c.v.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// Get 從設定檔根據key 抓值
func (c *Config) Get(key string) interface{} {
	return c.v.Get(key)
}
