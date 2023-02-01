package retry

import (
	"log"
)

// 一个重试框架的封装
func RetryFrame(x func() error, attempt int, s string) bool {
	retry := false
	var err error
	for i := 1; i <= attempt; i++ {
		// x()返回true, 请求重试
		if retry {
			log.Printf("[%s] %s, retrying\n", s, err)
		}
		if err = x(); err == nil {
			return true
		}
		retry = true
	}

	// failed
	return false
}
