// winmetric project winmetric.go
package win

import (
	"unicode/utf16"
)

// 将utf16数组转换为字符串数组，注意
// C中字符串以'\0'结尾，那么utf16数组最后2个元素有可能都是'\0’
// 转换为字符串数组后，字符串数组最后一个元素或者唯一的元素
// 可能是空字符串
func UTF16ToString(s []uint16) []string {
	var str []string = []string{}
	var zero int = 0
	for i, v := range s {
		if v == 0 {
			str = append(str, string(utf16.Decode(s[zero:i])))
			zero = i + 1
		}

	}
	return str
}
