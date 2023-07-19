package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"reflect"
)

func TryGet(key string, defaultValue interface{}) interface{} {
	value := viper.Get(key)
	if value == nil {
		logrus.Warnf("key[%s]未获取到nacos配置，使用默认值: %v", key, defaultValue)
		return defaultValue
	}

	defaultValueReflect := reflect.ValueOf(defaultValue)
	defaultValueType := defaultValueReflect.Type()

	valueReflect := reflect.ValueOf(value)
	valueType := valueReflect.Type()

	// 检查默认值类型和值类型是否匹配
	if defaultValueType != valueType {
		logrus.Warnf("key[%s]配置值类型不匹配，期望类型为 %s，实际类型为 %s",
			key, defaultValueType, valueType)
		return defaultValue
	}

	return value
}

func TryGetThrow(key string, defaultValue interface{}) interface{} {
	value := TryGet(key, defaultValue)
	if value == defaultValue {
		logrus.Fatalf("key[%s]未获取到nacos配置!", key)
	}
	return value
}
