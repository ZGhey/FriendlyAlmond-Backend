package utils

import (
	"fmt"
	"github.com/asim/go-micro/v3/util/file"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

func LoadConfigFile(filePath string) error {
	isExisted, err := file.Exists(filePath)
	if err != nil {
		return err
	}
	if !isExisted {
		return fmt.Errorf("file: '%s' is not existed. %v", filePath, err)
	}
	viper.SetConfigFile(filePath)
	viper.SetConfigType("yaml")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath("./conf/")
	//viper.AddConfigPath(".")
	//viper.AddConfigPath("../conf/")
	return viper.ReadInConfig()
}

func LoadConfigString(str string) error {
	viper.SetConfigType("yaml")
	return viper.ReadConfig(strings.NewReader(str))
}

func LoadConfigViper(filePath string) (*viper.Viper, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Println(absPath)
		return nil, err
	}
	fmt.Println(absPath)

	v := viper.New()
	isExisted, err := file.Exists(absPath)
	if err != nil {
		return nil, err
	}
	if !isExisted {
		return nil, fmt.Errorf("file: '%s' is not existed. %v", absPath, err)
	}
	v.SetConfigType("yaml")
	v.SetConfigFile(absPath)
	if err = v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}

func GetConfig(key string) string {
	return viper.GetString(key)
}
func GetConfigStr(key string) string {
	return viper.GetString(key)
}

func GetConfigInt(key string) int {
	return viper.GetInt(key)
}

func GetConfigInt64(key string) int64 {
	return viper.GetInt64(key)
}

func GetConfigUint(key string) uint64 {
	return viper.GetUint64(key)
}

func GetConfigFloat64(key string) float64 {
	return viper.GetFloat64(key)
}
