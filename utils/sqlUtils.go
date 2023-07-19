package utils

import (
	"reflect"
	"unicode"
)

// GetFieldName 返回字段的下划线命名

func WhereWrapper(field interface{}) string {
	return GetFieldName(field) + " = "
}
func GetFieldName(field interface{}) string {
	value := reflect.ValueOf(field)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	typeOfStruct := value.Type()

	for i := 0; i < typeOfStruct.NumField(); i++ {
		fieldName := typeOfStruct.Field(i).Name
		//这里可以自定义标签来实现强制绑定字段，也可以是用json
		tag := typeOfStruct.Field(i).Tag.Get("json")

		if tag == "-" {
			continue // 跳过不需要处理的字段
		}
		//
		//if tag != "" {
		//	如果有 JSON tag，则使用 JSON tag 的值作为字段名
		//return strings.Split(tag, ",")[0]
		//}
		// 将字段名转换为下划线命名
		return ToSnakeCase(fieldName)
	}

	return ""
}

// ToSnakeCase 将驼峰命名转换为下划线命名
func ToSnakeCase(s string) string {
	var result []rune

	for i, r := range s {
		if i > 0 && (unicode.IsUpper(r) || (unicode.IsDigit(r) && !unicode.IsDigit(rune(s[i-1])))) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}

	return string(result)
}
