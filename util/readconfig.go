package util

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func ReadConfigAndSetup(configName string, GConfig interface{}) error {
	v, err := ReadConfig(configName)
	if err != nil {
		return err
	}
	err = v.Unmarshal(GConfig)
	if err != nil {
		return err
	}
	return nil
}

func ReadConfigToBytes(configName string) ([]byte, error) {
	v, err := ReadConfig(configName)
	if err != nil {
		return nil, err
	}
	settings := v.AllSettings()
	bytes, err := yaml.Marshal(settings)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func ReadConfig(configName string) (*viper.Viper, error) {
	// 预加载环境变量
	viper.AutomaticEnv()
	// 获取一个viper实例
	vconfig := viper.New()
	// 设置配置文件的路径
	vconfig.SetConfigName(configName)
	vconfig.AddConfigPath(".")
	vconfig.SetConfigType("yaml")
	// 读取配置文件
	err := vconfig.ReadInConfig()
	if err != nil {
		_ = fmt.Errorf("service error: %s\n", err)
		return nil, err
	}

	//err = vconfig.Unmarshal(GConfig)
	//if err != nil {
	//	_ = fmt.Errorf("service error: %s\n", err)
	//	return err
	//}
	return vconfig, nil
}
