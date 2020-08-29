package acgo

import (
	"unicode"
)

// IsDigital 判断一个字符串是不是全是数字
func IsDigital(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}

	return true
}
