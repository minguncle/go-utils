package utils

import (
	"reflect"
)

// ObjectUtils 是一个用于操作对象的工具类
type ObjectUtils struct{}

// IsNull 检查对象是否为空（nil）或零值
func IsNull(obj interface{}) bool {
	if obj == nil {
		return true
	}

	value := reflect.ValueOf(obj)
	kind := value.Kind()

	switch kind {
	case reflect.String, reflect.Slice, reflect.Map, reflect.Array:
		return value.Len() == 0
	default:
		zeroValue := reflect.Zero(value.Type()).Interface()
		return reflect.DeepEqual(obj, zeroValue)
	}
}

// IfNullThrow 检查对象是否为空，如果为空则抛出指定错误
func IfNullThrow(obj interface{}, err error) error {
	if IsNull(obj) {
		return err
	}
	return nil
}
