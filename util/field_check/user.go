package fieldcheck

import (
	"net/url"
	"regexp"
	"time"
	"unicode"
)

// 用户名, 要求长度[4, 20], 只允许大小写字符, 数字以及下划线
func CheckUsernameValid(username string) bool {
	for _, v := range username {
		if !unicode.IsLetter(v) && !unicode.IsDigit(v) && v != '_' {
			return false
		}
	}

	l := len(username)
	return l >= 4 && l <= 20
}

// 密码, 要求长度[6, 25], 只允许大小写字符, 数字以及下划线
func CheckPasswordValid(password string) bool {
	for _, v := range password {
		if !unicode.IsLetter(v) && !unicode.IsDigit(v) && v != '_' {
			return false
		}
	}

	l := len(password)
	return l >= 6 && l <= 25
}

func CheckLinkVaild(link string) bool {
	_, err := url.ParseRequestURI(link)
	return err == nil
}

func CheckEmailVaild(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func CheckIntroductionVaild(s string) bool {
	l := len(s)
	return l >= 1 && l <= 300
}

func CheckPhoneVaild(s string) bool {
	reg := "^1[345789]{1}\\d{9}$"
	r := regexp.MustCompile(reg)
	return r.MatchString(s)
}

func CheckGenderVaild(g int) bool {
	return g == 1 || g == 0
}

func CheckBirthdayVaild(s string) bool {
	_, err := time.Parse("2006-01-02", s)
	return err == nil
}
