package fieldcheck

import "unicode/utf8"

func CheckLabelVaile(s string) bool {
	rcnt := utf8.RuneCountInString(s)
	// 1~20个utf8字符
	return rcnt >= 1 && rcnt <= 20
}
