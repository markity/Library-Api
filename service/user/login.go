package user

import (
	"database/sql"
	"library-api/dao"
)

func TryCheckLoginInfo(username string, passwordCrypto []byte) (bool, int64, error) {
	row := dao.DB.QueryRow("SELECT id FROM user WHERE username = ? AND password_crypto = ?",
		username, passwordCrypto)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, 0, nil
		} else {
			return false, 0, err
		}
	}

	return true, id, nil
}
