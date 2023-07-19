package utils

// Optional 是一个类似于 Java Optional 的类型
type Optional struct {
	value interface{}
}

// Of 创建一个包含指定值的 Optional
func Of(value interface{}) Optional {
	return Optional{value: value}
}

// IsPresent 检查 Optional 是否包含值
func (opt Optional) IsPresent() bool {
	return opt.value != nil
}

// IfPresent 如果 Optional 包含值，则执行指定的函数
func (opt Optional) IfPresent(fn func(interface{})) {
	if opt.value != nil {
		fn(opt.value)
	}
}

// OrElse 如果 Optional 包含值，则返回该值，否则返回指定的默认值
func (opt Optional) OrElse(defaultValue interface{}) interface{} {
	if opt.value != nil {
		return opt.value
	}
	return defaultValue
}

// OrElseThrow 如果 Optional 包含值，则返回该值，否则抛出指定的错误
func (opt Optional) OrElseThrow(err error) (interface{}, error) {
	if opt.value != nil {
		return opt.value, nil
	}
	return nil, err
}
