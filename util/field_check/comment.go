package fieldcheck

import "unicode/utf8"

func CheckCommentVaile(s string) bool {
	n := utf8.RuneCountInString(s)
	return n >= 10 && n <= 500
}
