package user

import (
	"library-api/dao"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

// 如果出错, 应当重试, 如果没有出错, 那么判断inserted的值
// 以此来判断该用户名是否已经被占用了
func TryCreateUser(username string, passwordCrypto []byte, createdAt time.Time) (err error, inserted bool) {
	_, err = dao.DB.Exec(`INSERT INTO user(username, password_crypto, created_at) VALUES(?,?,?)`, username,
		passwordCrypto, createdAt)
	if err != nil {
		mError := err.(*mysql.MySQLError)
		// 1062 duplicate entry, 代表该用户已存在
		if mError.Number == 1062 {
			return nil, false
		} else {
			log.Printf("insert error in TryCreateUser: %v\n", err)
			return err, false
		}
	}

	return nil, true
}
